package rpc

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm/clause"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"

	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
	"gorm.io/gorm"
)

type billConsumeRecord struct {
	ID         uuid.UUID
	GiftAmount int32
	BaseAmount int32
	Product    *fields.ProductPackageTicketArr
	Package    *fields.ProductPackageTicketArr
	OldAccount *model.TableCardAccount
	NewAccount *model.TableCardAccount
}

type consumeRequest struct {
	Tx         *gorm.DB    //事务
	CardID     uuid.UUID   //消费的卡ID,用于查询卡详情
	AccountIDs []uuid.UUID //可用于扣款的账户集合
	CostValue  int32
	Products   *fields.ProductPackageTicketArr
	Packages   *fields.ProductPackageTicketArr
	BranchID   uuid.UUID
	StaffID    uuid.UUID
	MerchantID uuid.UUID
	BillType   string //账单类型
	PosBillID  uuid.UUID
}

type updatePPRequest struct {
	Server         *Server
	Ctx            context.Context
	AccountBalance *fields.ProductPackageTicketArr       //流水所属账户的剩余赠品/套餐
	BillBalance    *fields.ProductPackageTicketArr       //流水的剩余赠品/套餐
	UpdatedBalance []*proto.CostProductPackageTicketItem //更新之后的
	BranchID       uuid.UUID
	Category       model.ProductPackageCategory //package or product
}

type updatePPResponse struct {
	AfterAccountBalance fields.ProductPackageTicketArr
	NewPPBill           []model.TableProductPackage
}

// GetAccounts 获取账户列表
func (s *Server) GetAccounts(ctx context.Context, request *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetAccounts")
	resp := &proto.GetAccountsResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)
	var branchIDs []uuid.UUID
	branchID := uuid.FromStringOrNil(request.BranchId)
	if branchID != uuid.Nil {
		branchIDs = append(branchIDs, branchID)
	} else {
		branchIDs = metadata.BranchIDs
	}

	count, err := model.CountAccounts(request.Status, branchIDs, metadata.MerchantID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("model.CountAccounts 账户查询数据库错误:%v", err))
		resp.ErrorMessage = "查询失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if count == 0 {
		resp.Data = &proto.AccountList{Total: 0}
		return resp, nil
	}

	accounts, err := model.GetAccounts(request.Status, branchIDs, request.Offset, request.Limit, metadata.MerchantID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAccounts 账户查询数据库错误:%v", err))
		resp.ErrorMessage = "查询失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var accountData []*proto.Account

	for _, account := range accounts {
		accountData = append(accountData, toAccountProto(&account))
	}
	resp.Data = &proto.AccountList{Accounts: accountData, Total: int32(count)}
	return resp, nil
}

// ShowAccount 查询单条账户信息
func (s *Server) ShowAccount(ctx context.Context, request *proto.ShowAccountRequest) (*proto.ShowAccountResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowAccount")
	resp := &proto.ShowAccountResponse{
		ErrorCode: pkgs.Success,
	}
	accountID := uuid.FromStringOrNil(request.Id)
	if accountID == uuid.Nil {
		util.Logger.Error(fmt.Sprintf("ShowAccount 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	account, err := model.ShowCardWithFreeze(accountID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAccounts 账户查询数据库错误:%v", err))
		resp.ErrorMessage = "查询失败"
		resp.ErrorCode = pkgs.ErrInternal
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "未找到所查询的账户"
		}
		return resp, nil
	}
	resp.Data = toAccountWithFreezeProto(account)
	return resp, nil
}

// GetAccountsByCardID 根据卡id查询账户信息
func (s *Server) GetAccountsByCardID(ctx context.Context, request *proto.GetAccountsByCardIDRequest) (*proto.GetAccountsByCardIDResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetAccountsByCardID")
	resp := &proto.GetAccountsByCardIDResponse{
		ErrorCode: pkgs.Success,
	}

	cardID := uuid.FromStringOrNil(request.Id)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	card, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "卡不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("ShowAccountByCardID 账户查询数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询账户信息失败"
		return resp, nil
	}

	// 如果是副卡，则需要查询主卡账户信息
	if card.PrimaryID != nil {
		cardID = *card.PrimaryID
	}

	accounts, err := model.GetCardAccountsByCardID(cardID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "账户不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("ShowAccountByCardID 账户查询数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询账户信息失败"
		return resp, nil
	}

	for i := 0; i < len(accounts); i++ {
		accountTag := &proto.AccountWithTag{Account: toAccountProto(&accounts[i])}
		tagResp, err := s.merchantBasic().ShowBranchTag(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.ShowBranchTagRequest{Id: accounts[i].TagID.String()})
		if err != nil || tagResp.ErrorCode != pkgs.Success {
			util.Logger.Error(fmt.Sprintf("ShowAccountByCardID 查询门店标签错误,err:%v, resp:%v", err, tagResp))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "查询账户信息失败"
			return resp, nil
		}
		accountTag.TagName = tagResp.Data.Name
		for _, v := range tagResp.Data.Branches {
			accountTag.BranchIds = append(accountTag.BranchIds, v)
		}
		resp.Data = append(resp.Data, accountTag)
	}
	return resp, nil
}

// UpdateAccountStatus 冻结/解冻
func (s *Server) UpdateAccountStatus(ctx context.Context, request *proto.UpdateAccountStatusRequest) (*proto.NoDataResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateAccountStatus")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}
	accountID := uuid.FromStringOrNil(request.Id)
	metaData := pkgs.GetMetadata(ctx)
	//todo 冻结原因
	if accountID == uuid.Nil || len(request.Reason) == 0 {
		util.Logger.Error(fmt.Sprintf("UpdateAccountStatus 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	account, err := model.ShowCardAccount(accountID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("ShowCardAccount 查询账户数据库错误:%+v", err))
		resp.ErrorMessage = "账户数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "账户不存在"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}

	newStatus, err := validAccountStatus(model.AccountStatus(account.Status), request.Action)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("冻结解冻账户失败， 账户状态异常: status:%s, action:%s", account.Status, request.Action))
		resp.ErrorMessage = "账户状态异常"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	updateAccount := &model.TableCardAccount{
		ID:     account.ID,
		Status: newStatus,
	}

	accountFreeze := &model.TableAccountFreeze{
		ID:         uuid.NewV4(),
		AccountID:  &accountID,
		Action:     request.Action,
		Reason:     request.Reason,
		StaffID:    &metaData.StaffID,
		MerchantID: &metaData.MerchantID,
	}

	if err := model.UpdateAccountStatusByID(updateAccount, accountFreeze); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAccountByID 数据库错误:%+v", err))
		resp.ErrorMessage = "更新出错了"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	return resp, nil
}

// GetAccountByMemberID 会员ID查账户信息
func (s *Server) GetAccountByMemberID(ctx context.Context, request *proto.GetAccountByMemberIDRequest) (*proto.GetMemberAccountResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetAccountByMemberID")
	resp := &proto.GetMemberAccountResponse{
		ErrorCode: pkgs.Success,
	}

	memberID := uuid.FromStringOrNil(request.MemberId)
	if memberID == uuid.Nil {
		util.Logger.Error(fmt.Sprintf("GetAccountByMemberID 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	cards, err := model.GetMemberCardsByMemberID(memberID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetMemberCardsByMemberID 数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	cardWithAccountSlice, err := buildMemberAccounts(cards)
	if err != nil {
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = cardWithAccountSlice
	return resp, nil
}

// GetAccountByCardCode 刷卡查询卡账户信息
func (s *Server) GetAccountByCardCode(ctx context.Context, request *proto.GetAccountByCardCodeRequest) (*proto.GetAccountByCardCodeResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBill")
	resp := &proto.GetAccountByCardCodeResponse{
		ErrorCode: pkgs.Success,
	}

	metadata := pkgs.GetMetadata(ctx)
	cardCode := request.CardCode
	if cardCode == "" {
		util.Logger.Error(fmt.Sprintf("GetAccountByCardCode 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	card, err := model.ShowCardByCode(cardCode, metadata.MerchantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("ShowMemberCardsByCode 数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	var cardWithAccount proto.CardWithAccount
	accounts, err := model.GetCardAccountsByCardID(card.ID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetCardAccountsByCardID 数据库错误:%v", err))
		return nil, err
	}
	var accountsProto []*proto.Account
	if len(accounts) > 0 {
		for _, account := range accounts {
			accountsProto = append(accountsProto, toAccountProto(&account))
		}
	}
	cardWithAccount.Card = toCardProto(card)
	cardWithAccount.Accounts = accountsProto

	resp.Data = &cardWithAccount
	return resp, nil
}

// GetMemberAccounts 会员账户
func (s *Server) GetMemberAccounts(ctx context.Context, request *proto.GetMemberAccountsRequest) (*proto.GetMemberAccountsResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetMemberAccounts")
	resp := &proto.GetMemberAccountsResponse{
		ErrorCode: pkgs.Success,
	}
	memberID := uuid.FromStringOrNil(request.MemberId)
	if memberID == uuid.Nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}
	accounts, err := model.GetAccountByMemberID(memberID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAccountByMemberID 数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	if len(accounts) > 0 {
		var accountsProto []*proto.Account
		for _, a := range accounts {
			accountsProto = append(accountsProto, toAccountProto(&a))
		}
		resp.Data = accountsProto
	}

	return resp, nil
}

func buildMemberAccounts(cards *[]model.TableCard) ([]*proto.CardWithAccount, error) {
	var cardWithAccountSlice []*proto.CardWithAccount
	if len(*cards) > 0 {
		var accountIDs []uuid.UUID
		for _, card := range *cards {
			ids := card.AccountIDs
			if ids != nil {
				for _, id := range *ids {
					accountIDs = append(accountIDs, id)
				}
			}
		}
		accounts, err := model.GetCardAccountsByIDs(accountIDs)
		if err != nil {
			util.Logger.Error(fmt.Sprintf("GetCardAccountsByIDs 数据库错误:%v", err))
			return nil, err
		}
		accountMap := make(map[uuid.UUID]model.TableCardAccount)
		for _, account := range accounts {
			accountMap[account.ID] = account
		}

		//开始组装数据
		for _, card := range *cards {
			var cardWithAccount proto.CardWithAccount
			cardWithAccount.Card = toCardProto(&card)
			var accountsProto []*proto.Account
			if card.AccountIDs != nil {
				for _, id := range *card.AccountIDs {
					if account, ok := accountMap[id]; ok {
						accountsProto = append(accountsProto, toAccountProto(&account))
					}
				}
				cardWithAccount.Accounts = accountsProto
			}
			cardWithAccountSlice = append(cardWithAccountSlice, &cardWithAccount)
		}
	}
	return cardWithAccountSlice, nil
}

// UpdateAccountBalance 修改充值流水余额
func (s *Server) UpdateAccountBalance(ctx context.Context, request *proto.UpdateAccountBalanceRequest) (*proto.NoDataResponse, error) {
	operatorComment := "修改余额"
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateAccountBalance")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)
	branchID := uuid.FromStringOrNil(request.BranchId)
	if request.BillCode == "" || branchID == uuid.Nil || len(request.Reason) == 0 {
		util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	//当前账单信息
	billInfo, err := model.ShowBillByCode(request.BillCode)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("model.ShowCard数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorMessage = "请求参数错误，账单不存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
		}
		return resp, nil
	}

	tx := model.GetDBEntity().Conn.Begin()
	var accountInfo model.TableCardAccount
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", billInfo.AccountID).First(&accountInfo).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBalance ShowCard数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//计算金额差值
	baseValueDiff := request.BaseValue - billInfo.BaseValueLeft
	giftValueDiff := request.GiftValue - billInfo.GiftValueLeft

	//原有流水的商品、套餐
	var oldProducts fields.ProductPackageTicketArr
	var oldPackages fields.ProductPackageTicketArr

	if len(billInfo.ProductPackageBills) > 0 {
		for _, bill := range billInfo.ProductPackageBills {
			if bill.Left > 0 {
				item := fields.ProductPackageTicket{
					ID:    bill.ProductPackageID,
					Code:  bill.Code,
					Count: bill.Left,
					Price: bill.Price,
					Title: bill.Title,
				}
				switch bill.Category {
				case model.ProductPackageCategoryProduct:
					oldProducts = append(oldProducts, item)
				case model.ProductPackageCategoryPackage:
					oldPackages = append(oldPackages, item)
				}
			}
		}
	}

	//组装变更前balance
	oldPackagesParams := pkgs.MakeParamsArr(oldPackages)
	oldProductsParams := pkgs.MakeParamsArr(oldProducts)
	oldAccountBillBalance := model.AccountBalanceBill{
		BaseValue: accountInfo.BaseValue,
		GiftValue: accountInfo.GiftValue,
		Packages:  &oldPackagesParams,
		Products:  &oldProductsParams,
	}
	//组装变更后balance
	newPackagesParams := pkgs.MakeParamsArr(request.Packages)
	newProductsParams := pkgs.MakeParamsArr(request.Products)
	newAccountBillBalance := model.AccountBalanceBill{
		BaseValue: request.BaseValue,
		GiftValue: request.GiftValue,
		Packages:  &newPackagesParams,
		Products:  &newProductsParams,
	}

	updateAccount := make(map[string]interface{})
	afterAccount := accountInfo
	updateAccount["base_value"] = accountInfo.BaseValue - billInfo.BaseValueLeft + request.BaseValue
	updateAccount["gift_value"] = accountInfo.GiftValue - billInfo.GiftValueLeft + request.GiftValue
	afterAccount.BaseValue = accountInfo.BaseValue - billInfo.BaseValueLeft + request.BaseValue
	afterAccount.GiftValue = accountInfo.GiftValue - billInfo.GiftValueLeft + request.GiftValue
	productRequest := updatePPRequest{
		Server:         s,
		Ctx:            ctx,
		AccountBalance: accountInfo.Products,
		BillBalance:    &oldProducts,
		UpdatedBalance: request.Products,
		BranchID:       branchID,
		Category:       model.ProductPackageCategoryProduct,
	}
	productResponse, err := buildNewProductPackageBalance(productRequest)
	if err != nil {
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	packageRequest := updatePPRequest{
		Server:         s,
		Ctx:            ctx,
		AccountBalance: accountInfo.Packages,
		BillBalance:    &oldPackages,
		UpdatedBalance: request.Packages,
		BranchID:       branchID,
		Category:       model.ProductPackageCategoryPackage,
	}
	packageResponse, err := buildNewProductPackageBalance(packageRequest)
	if err != nil {
		tx.Rollback()
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	updateAccount["products"] = productResponse.AfterAccountBalance
	updateAccount["packages"] = packageResponse.AfterAccountBalance
	afterAccount.Products = &productResponse.AfterAccountBalance
	afterAccount.Packages = &packageResponse.AfterAccountBalance

	//①账户流水表插入新数据 更新账户表数据
	dealMap := make(map[model.BillType]int32) //消费/充值 -> 金额
	updateBillValueMap := make(map[string]int32, 2)
	if baseValueDiff >= 0 && giftValueDiff >= 0 {
		//一笔，充值
		dealMap[model.BillTypeRecharge] = baseValueDiff + giftValueDiff
	} else if baseValueDiff < 0 && giftValueDiff < 0 {
		//消费，一笔
		baseValueDiff = int32(math.Abs(float64(baseValueDiff)))
		giftValueDiff = int32(math.Abs(float64(giftValueDiff)))
		dealMap[model.BillTypeConsume] = baseValueDiff + giftValueDiff
		updateBillValueMap["base_value_left"] = request.BaseValue
		updateBillValueMap["gift_value_left"] = request.GiftValue

	} else {
		//分开计算 两笔
		//本金充值 赠金消费
		if baseValueDiff >= 0 {
			giftValueDiff = int32(math.Abs(float64(giftValueDiff)))
			dealMap[model.BillTypeRecharge] = baseValueDiff
			dealMap[model.BillTypeConsume] = giftValueDiff
			updateBillValueMap["gift_value_left"] = request.GiftValue
		} else {
			//本金消费 赠金充值
			baseValueDiff = int32(math.Abs(float64(baseValueDiff)))
			dealMap[model.BillTypeRecharge] = giftValueDiff
			dealMap[model.BillTypeConsume] = baseValueDiff
			updateBillValueMap["base_value_left"] = request.BaseValue
		}
	}
	var newBills []model.TableAccountBill
	newRechargeBillID := billInfo.ID //新的充值记录的ID，用于后面pp流水的记录，默认是原来充值流水的ID
	for billType, value := range dealMap {
		if value > 0 {
			billCode, err := s.getBillCode(ctx, branchID, time.Now())
			if err != nil {
				tx.Rollback()
				util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 生成billCode错误:%+v", err))
				resp.ErrorMessage = "数据库错误"
				resp.ErrorCode = pkgs.ErrInternal
				return resp, nil
			}
			bill := model.TableAccountBill{
				ID:              uuid.NewV4(),
				BillCode:        billCode,
				AccountID:       billInfo.AccountID,
				CardID:          billInfo.CardID,
				CardCode:        billInfo.CardCode,
				BranchID:        &branchID,
				ChangeValue:     baseValueDiff + giftValueDiff,
				ChangeCategory:  model.BillCategoryChange,
				ChangeType:      billType,
				BaseValue:       baseValueDiff,
				GiftValue:       giftValueDiff,
				AfterAccount:    &afterAccount,
				StaffID:         &metadata.StaffID,
				OperatorComment: operatorComment,
				MerchantID:      &metadata.MerchantID,
			}
			//充值，要给余额
			if billType == model.BillTypeRecharge {
				bill.BaseValueLeft = baseValueDiff
				bill.GiftValueLeft = giftValueDiff
				newRechargeBillID = bill.ID
			}
			newBills = append(newBills, bill)
		}
	}
	//流水表插入新记录
	err = tx.Create(&newBills).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 插入新流水数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	//更新账户余额
	err = tx.Model(model.TableCardAccount{}).Where("id = ?", accountInfo.ID).Updates(updateAccount).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 更新账户数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//②如果delMap存在消费,更新原bill表余额
	if len(updateBillValueMap) > 0 {
		err = tx.Model(model.TableAccountBill{}).Where("id = ?", billInfo.ID).Updates(updateBillValueMap).Error
		if err != nil {
			tx.Rollback()
			util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 更新原流水剩余金额数据库错误:%+v", err))
			resp.ErrorMessage = "数据库错误"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	}
	//③修改商品套餐水流余额置零
	err = tx.Model(model.TableProductPackage{}).Where("bill_id = ? AND category = ?", billInfo.ID, model.ProductPackageCategoryPackage).Update("left", 0).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBillBalance 更新原流水剩余套餐数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	err = tx.Model(model.TableProductPackage{}).Where("bill_id = ? AND category = ?", billInfo.ID, model.ProductPackageCategoryProduct).Update("left", 0).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBillBalance 更新原流水剩余商品数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//④ 插入新的商品流水数据
	var newPPBills []model.TableProductPackage
	for _, bill := range productResponse.NewPPBill {
		bill.BillID = &newRechargeBillID
		newPPBills = append(newPPBills, bill)
	}
	for _, bill := range packageResponse.NewPPBill {
		bill.BillID = &newRechargeBillID
		newPPBills = append(newPPBills, bill)
	}

	err = tx.Create(&newPPBills).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 添加新商品流水数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	//⑤组装记录表数据
	createUpdateAccountBill := model.TableUpdateAccountBill{
		ID:                    uuid.NewV4(),
		AccountBillID:         &billInfo.ID,
		OldAccountBillBalance: &oldAccountBillBalance,
		NewAccountBillBalance: &newAccountBillBalance,
		BranchID:              &branchID,
		StaffID:               &metadata.StaffID,
		MerchantID:            &metadata.MerchantID,
		Reason:                request.Reason,
	}

	err = tx.Create(&createUpdateAccountBill).Error
	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("UpdateAccountBalance 插入更新记录数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	tx.Commit()
	return resp, nil
}

//处理商品套餐的变更计算
func buildNewProductPackageBalance(req updatePPRequest) (*updatePPResponse, error) {
	oldMap := make(map[uuid.UUID]*fields.ProductPackageTicket)
	newPackageProductBillsMap := make(map[uuid.UUID]*model.TableProductPackage)
	if req.AccountBalance != nil {
		for _, item := range *req.AccountBalance {
			oldMap[item.ID] = &item
		}
		//流水里的商品全部减掉
		if req.BillBalance != nil {
			for _, item := range *req.BillBalance {
				//肯定存在
				oldMap[item.ID].Count -= item.Count
			}
		}
	}

	if len(req.UpdatedBalance) > 0 {
		var newPPIDs fields.UUIDArr
		for _, item := range req.UpdatedBalance {
			id := uuid.FromStringOrNil(item.Id)
			newBill := model.TableProductPackage{
				ID:               uuid.NewV4(),
				ProductPackageID: id,
				Number:           item.Number,
				Left:             item.Number,
				Category:         req.Category,
			}
			if ii, ok := oldMap[id]; ok {
				//原来就有，直接加
				oldMap[id].Count += item.Number
				newBill.Code = ii.Code
				newBill.Title = ii.Title
				newBill.Price = ii.Price
			} else {
				//没有，需要查询详情
				newPPIDs = append(newPPIDs, id)
				oldMap[id] = &fields.ProductPackageTicket{
					ID:    id,
					Count: item.Number,
				}
			}
			newPackageProductBillsMap[id] = &newBill
		}

		if len(newPPIDs) > 0 {
			request := merchantBasic.MultiGetGoodsAndPackagesRequest{BranchId: req.BranchID.String()}
			if req.Category == model.ProductPackageCategoryProduct {
				request.GoodsIds = newPPIDs.ToStringArr()
			} else {
				request.PackageIds = newPPIDs.ToStringArr()
			}
			resp, err := req.Server.merchantBasic().MultiGetGoodsAndPackages(pkgs.MetadataContent(pkgs.GetMetadata(req.Ctx)), &request)
			if err != nil || resp == nil || resp.ErrorCode != pkgs.Success {
				util.Logger.Error(fmt.Sprintf("UpdateBillBalance 查询商品详情错误:%v, %v", err, resp))
				return nil, errors.New("查询商品详情错误")
			}
			if req.Category == model.ProductPackageCategoryProduct {
				for _, pp := range resp.Data.Goods {
					id := uuid.FromStringOrNil(pp.Id)
					oldMap[id].Code = pp.Code
					oldMap[id].Price = pp.Price
					oldMap[id].Title = pp.Name
					if _, ok := newPackageProductBillsMap[id]; ok {
						newPackageProductBillsMap[id].Code = pp.Code
						newPackageProductBillsMap[id].Price = pp.Price
						newPackageProductBillsMap[id].Title = pp.Name
					}
				}
			} else {
				for _, pp := range resp.Data.Packages {
					id := uuid.FromStringOrNil(pp.Id)
					oldMap[id].Code = pp.Code
					oldMap[id].Price = pp.Price
					oldMap[id].Title = pp.Name
					if _, ok := newPackageProductBillsMap[id]; ok {
						newPackageProductBillsMap[id].Code = pp.Code
						newPackageProductBillsMap[id].Price = pp.Price
						newPackageProductBillsMap[id].Title = pp.Name
					}
				}
			}
		}
	}

	var resAccount fields.ProductPackageTicketArr
	var resBills []model.TableProductPackage
	for _, item := range oldMap {
		//过滤余额为0的商品套餐
		if item.Count > 0 {
			pp := *item
			resAccount = append(resAccount, pp)
		}
	}
	for _, item := range newPackageProductBillsMap {
		resBills = append(resBills, *item)
	}
	return &updatePPResponse{
		AfterAccountBalance: resAccount,
		NewPPBill:           resBills,
	}, nil

}

// AddAccountDeduction 增加扣款
func (s *Server) AddAccountDeduction(ctx context.Context, request *proto.AddAccountDeductionRequest) (*proto.NoDataResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("AddAccountDeduction")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}

	metadata := pkgs.GetMetadata(ctx)
	branchID := uuid.FromStringOrNil(request.BranchId)
	cardID := uuid.FromStringOrNil(request.CardId)
	staffID := metadata.StaffID
	merchantID := metadata.MerchantID

	//不用前端上传的account_ids,根据branch_id查询可用账户
	protoResp, err := s.merchantBasic().GetBranchTags(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.GetBranchTagsRequest{
		BranchIds: []string{request.BranchId},
	})

	if err != nil || protoResp.ErrorCode != pkgs.Success {
		util.Logger.Error(fmt.Sprintf("GetBranchTags 获取门店标签错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "门店标签查询错误"
		return resp, nil
	}

	if protoResp.Data.Total == 0 {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "门店无可用标签"
		return resp, nil
	}
	var tagIDs []interface{}
	for _, tag := range protoResp.Data.BranchTags {
		tagIDs = append(tagIDs, tag.Id)
	}
	tx := model.GetDBEntity().Conn.Begin()
	var accountIDs []uuid.UUID
	err = tx.Model(&model.TableCardAccount{}).Select("id").Scopes(util.ColumnInScope("tag_id", tagIDs)).Scan(&accountIDs).Error

	if err != nil {
		util.Logger.Error(fmt.Sprintf("查询可用账户 数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "查询可用账户数据库错误"
		return resp, nil
	}

	consumeRequest := &consumeRequest{
		Tx:         tx,
		CardID:     cardID,
		AccountIDs: accountIDs,
		CostValue:  request.CostValue,
		BranchID:   branchID,
		StaffID:    staffID,
		MerchantID: merchantID,
		BillType:   string(model.BillTypeDeduction),
	}
	consumeBillsMap, iErr := s.consume(ctx, consumeRequest)
	if iErr.ErrorCode != 0 {
		tx.Rollback()
		resp.ErrorMessage = iErr.ErrorMessage
		resp.ErrorCode = iErr.ErrorCode
		return resp, nil
	}

	//addDeduction表添加记录
	var consumeBillIDs fields.UUIDArr
	for _, bill := range consumeBillsMap {
		consumeBillIDs = append(consumeBillIDs, bill.ID)
	}
	createAddDeduction := &model.TableAddAccountDeduction{
		ID:         uuid.NewV4(),
		ConsumeIDs: &consumeBillIDs,
		BillNumber: request.BillNumber,
		BranchID:   &branchID,
		StaffID:    &staffID,
		MerchantID: &merchantID,
		Reason:     request.Reason,
	}

	err = model.AddAccountDeduction(tx, createAddDeduction)

	if err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("AddAccountDeduction 数据库错误:%+v", request))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	tx.Commit()
	return resp, nil
}

// TransferCardBalance 划账
func (s *Server) TransferCardBalance(ctx context.Context, request *proto.TransferCardBalanceRequest) (*proto.NoDataResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("TransferCardBalance")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)

	sourceAccountID := uuid.FromStringOrNil(request.SourceAccountId)
	destCardID := uuid.FromStringOrNil(request.DestCardId)
	operateBranchID := uuid.FromStringOrNil(request.OperateBranchId)
	merchantID := metadata.MerchantID
	if sourceAccountID == uuid.Nil || destCardID == uuid.Nil || request.Amount <= 0 || operateBranchID == uuid.Nil {
		util.Logger.Error(fmt.Sprintf("TransferCardBalance 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	sourceAccount, err := model.ShowCardAccount(sourceAccountID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("model.ShowCardAccount 数据库查询错误:%v", err))
		resp.ErrorMessage = "数据库查询错误"
		resp.ErrorCode = pkgs.ErrInternal
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "【转出】未查询到账户信息"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}
	if err = validateTransferAccount(sourceAccount); err != nil {
		util.Logger.Error(fmt.Sprintf("划卡 sourceAccount验证失败:%v", err))
		resp.ErrorMessage = "【转出】" + err.Error()
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	if request.Amount > sourceAccount.BaseValue+sourceAccount.GiftValue {
		util.Logger.Error("TransferCardBalance 卡余额不足")
		resp.ErrorMessage = "【转出】账户余额不足"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	cardInfo, destAccount, iErr := getTransferDestAccount(sourceAccount, destCardID, merchantID)
	if iErr.ErrorCode != 0 {
		util.Logger.Error(fmt.Sprintf("查询转账入账account失败:%v", iErr))
		resp.ErrorMessage = iErr.ErrorMessage
		resp.ErrorCode = iErr.ErrorCode
		return resp, nil
	}
	//无可用账户，新建
	if destAccount == nil {
		destAccount = &model.TableCardAccount{
			ID:         uuid.NewV4(),
			MemberID:   cardInfo.MemberID,
			BranchID:   &operateBranchID,
			Status:     model.AccountStatusActivated,
			MerchantID: merchantID,
			Category:   cardInfo.SubCategory, //账户类型跟卡类型一致
		}
	}
	//旧账户消费
	tx := model.GetDBEntity().Conn.Begin()
	req := &consumeRequest{
		Tx:         tx,
		AccountIDs: []uuid.UUID{sourceAccountID},
		CostValue:  request.Amount,
		BranchID:   operateBranchID,
		StaffID:    metadata.StaffID,
		MerchantID: merchantID,
		CardID:     cardInfo.ID,
		BillType:   string(model.BillTypeTransfer),
	}
	consumeResp, iErr := s.consume(ctx, req)
	if iErr.ErrorCode != 0 {
		resp.ErrorCode = iErr.ErrorCode
		resp.ErrorMessage = iErr.ErrorMessage
		return resp, nil
	}

	//转入账户添加充值流水
	billCode, err := s.getBillCode(ctx, operateBranchID, time.Now())
	if err != nil {
		util.Logger.Error(fmt.Sprintf("transfer 生成billCode失败:%v", err))
		resp.ErrorMessage = "生成billCode失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var changeBaseValue, changeGiftValue int32
	for _, bill := range consumeResp {
		changeBaseValue += bill.BaseValue
		changeGiftValue += bill.GiftValue
	}
	destAccount.BaseValue += changeBaseValue
	destAccount.GiftValue += changeGiftValue
	destAccount.TagID = sourceAccount.TagID

	if !util.ArrContainElement(cardInfo.AccountIDs.Slice(), destAccount.ID) {
		if cardInfo.AccountIDs == nil {
			cardInfo.AccountIDs = &fields.UUIDArr{}
		}
		*cardInfo.AccountIDs = append(*cardInfo.AccountIDs, destAccount.ID)
	}

	createRechargeBill := &model.TableAccountBill{
		ID:             uuid.NewV4(),
		BillCode:       billCode,
		AccountID:      &destAccount.ID,
		CardID:         &destCardID,
		BranchID:       &operateBranchID,
		ChangeValue:    request.Amount,
		ChangeCategory: model.BillCategoryRecharge,
		ChangeType:     model.BillTypeTransfer,
		BaseValue:      changeBaseValue,
		GiftValue:      changeGiftValue,
		AfterAccount:   destAccount,
		StaffID:        &metadata.StaffID,
		MerchantID:     &merchantID,
		BaseValueLeft:  changeBaseValue,
		GiftValueLeft:  changeGiftValue,
	}
	//划账表添加记录
	cardTransfer := &model.TableCardTransfer{
		ID:              uuid.NewV4(),
		SourceAccountID: &sourceAccountID,
		DestAccountID:   &destAccount.ID,
		TransferValue:   request.Amount,
		StaffID:         &metadata.StaffID,
		MerchantID:      &metadata.MerchantID,
	}

	if err := tx.Create(createRechargeBill).Error; err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("transfer 创建充值流水数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if err := tx.Save(destAccount).Error; err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("transfer 更新账户数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if err := tx.Updates(cardInfo).Error; err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("transfer 更新卡信息数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if err := tx.Create(cardTransfer).Error; err != nil {
		tx.Rollback()
		util.Logger.Error(fmt.Sprintf("transfer 划账记录数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	tx.Commit()

	return resp, nil
}

func toAccountProto(account *model.TableCardAccount) *proto.Account {
	var tagID string
	var products, packages []*proto.ProductPackage
	if account.TagID != nil {
		tagID = account.TagID.String()
	}
	if account.Products != nil {
		for _, v := range *account.Products {
			products = append(products, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Title:  v.Title,
			})
		}
	}
	if account.Packages != nil {
		for _, v := range *account.Packages {
			packages = append(packages, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Title:  v.Title,
			})
		}
	}

	return &proto.Account{
		Id:        account.ID.String(),
		MemberId:  util.UUIDToString(account.MemberID),
		BranchId:  util.UUIDToString(account.BranchID),
		BaseValue: account.BaseValue,
		GiftValue: account.GiftValue,
		Products:  products,
		Packages:  packages,
		TagId:     tagID,
		Status:    string(account.Status),
		CreateAt:  util.TimeUnix32(account.CreatedAt),
		Category:  string(account.Category),
	}
}

func toAccountWithFreezeProto(accountWithFreeze *model.TableCardAccount) *proto.AccountWithActionReason {
	var products, packages []*proto.ProductPackage
	var tagID string
	if accountWithFreeze.TagID != nil {
		tagID = accountWithFreeze.TagID.String()
	}
	if accountWithFreeze.Products != nil {
		for _, v := range *accountWithFreeze.Products {
			products = append(products, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Title:  v.Title,
			})
		}
	}
	if accountWithFreeze.Packages != nil {
		for _, v := range *accountWithFreeze.Packages {
			packages = append(packages, &proto.ProductPackage{
				Id:     v.ID.String(),
				Code:   v.Code,
				Number: v.Count,
				Title:  v.Title,
			})
		}
	}

	return &proto.AccountWithActionReason{
		Id:           accountWithFreeze.ID.String(),
		MemberId:     util.UUIDToString(accountWithFreeze.MemberID),
		BranchId:     util.UUIDToString(accountWithFreeze.BranchID),
		BaseValue:    accountWithFreeze.BaseValue,
		GiftValue:    accountWithFreeze.GiftValue,
		Products:     products,
		Packages:     packages,
		TagId:        tagID,
		Status:       string(accountWithFreeze.Status),
		CreateAt:     int32(accountWithFreeze.CreatedAt.Unix()),
		ActionReason: accountWithFreeze.FreezeInfo.Reason,
		Category:     string(accountWithFreeze.Category),
	}
}

// 验证账户状态
func validAccountStatus(currStatus model.AccountStatus, action string) (model.AccountStatus, error) {
	err := errors.New("账户状态异常")
	switch action {
	//冻结
	case model.AccountActionFreeze:
		if currStatus == model.AccountStatusActivated {
			return model.AccountStatusFrozen, nil
		}
		return "", err
	//解冻
	case model.AccountActionUnfreeze:
		if currStatus == model.AccountStatusFrozen {
			return model.AccountStatusActivated, nil
		}
		return "", err
	default:
		return "", err
	}
}

//Consume 通用消费
func (s *Server) consume(ctx context.Context, request *consumeRequest) (map[uuid.UUID]*model.TableConsumeBills, internalErr) {
	var iErr internalErr
	//卡信息，用于填充数据
	var cardInfo model.TableCard
	if err := request.Tx.Model(&model.TableCard{}).
		Where("id = ? and merchant_id = ?", request.CardID, request.MerchantID).
		First(&cardInfo).Error; err != nil {
		util.Logger.Error(fmt.Sprintf("consume 查询卡信息数据库错误:%v", err))
		iErr.ErrorMessage = "数据库查询失败"
		iErr.ErrorCode = pkgs.ErrInternal
		return nil, iErr
	}

	createConsumeBillsMap := make(map[uuid.UUID]*model.TableConsumeBills) //消费流水，对应到充值流水 key是消耗的充值流水ID
	updateBillsMap := make(map[uuid.UUID]map[string]interface{})          //需要更新的充值流水

	var createBills4Consume []*model.TableAccountBill //流水表消费类型数据，按账户ID分多条
	var updateAccounts []map[string]interface{}       //需要更新的账户列表
	var productNoLeftBillIDs []uuid.UUID              //库存置为0的商品流水
	var productUpdateBills []map[string]interface{}   //改库存的商品流水  id和left

	billAccounts4Consume := make(map[uuid.UUID]*billConsumeRecord) //账户id => 消费流水记录  中间变量 不参与到数据库处理

	dealAccounts := make(map[uuid.UUID]model.TableCardAccount) //涉及到的账户

	//余额
	if request.CostValue > 0 {
		//可用充值流水
		//开始扣费
		consumeAmountLeft := request.CostValue //需要扣的金额，每次循环后变小
		var usableValueBills []model.TableAccountBill
		if err := request.Tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("account_id in (?) AND (base_value_left > 0 or gift_value_left > 0)", request.AccountIDs).
			Find(&usableValueBills).Error; err != nil {
			util.Logger.Error(fmt.Sprintf("consume 查询流水数据库错误:%+v", err))
			iErr.ErrorMessage = "数据库错误"
			iErr.ErrorCode = pkgs.ErrInternal
			return nil, iErr
		}

		//涉及到的account
		var accounts []model.TableCardAccount
		if err := request.Tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id in (?)", request.AccountIDs).
			Find(&accounts).Error; err != nil {
			util.Logger.Error(fmt.Sprintf("consume 查询涉及账户数据库错误:%+v", err))
			iErr.ErrorMessage = "数据库错误"
			iErr.ErrorCode = pkgs.ErrInternal
			return nil, iErr
		}

		for _, account := range accounts {
			dealAccounts[account.ID] = account
		}

		//开始扣钱,先判断赠金够不够
		var giftValueTotal, baseValueTotal int32
		for _, bill := range usableValueBills {
			giftValueTotal += bill.GiftValueLeft
			baseValueTotal += bill.BaseValueLeft
		}

		//余额不足
		if request.CostValue > (giftValueTotal + baseValueTotal) {
			iErr.ErrorMessage = "账户余额不足"
			iErr.ErrorCode = pkgs.ErrUnprocessableEntity
			return nil, iErr
		}

		breakFlag := false

		//处理赠金
		for _, bill := range usableValueBills {
			if _, ok := billAccounts4Consume[*bill.AccountID]; !ok {
				accountInfo := dealAccounts[*bill.AccountID]
				if accountInfo.Status != model.AccountStatusActivated {
					util.Logger.Error(fmt.Sprintf("AddAccountDeduction 账户状态异常:%v", accountInfo))
					iErr.ErrorMessage = "账户状态异常"
					iErr.ErrorCode = pkgs.ErrUnprocessableEntity
					return nil, iErr
				}
				newAccount := new(model.TableCardAccount)
				*newAccount = accountInfo
				billAccounts4Consume[*bill.AccountID] = &billConsumeRecord{
					ID:         uuid.NewV4(),
					OldAccount: &accountInfo,
					NewAccount: newAccount,
				}
			}

			currRecord := billAccounts4Consume[*bill.AccountID] //肯定存在

			if bill.GiftValueLeft > 0 {
				//消费流水记录，可能还要更新baseValue，所以用map
				consumeBill := &model.TableConsumeBills{
					ID:            uuid.NewV4(),
					BillID:        &currRecord.ID, //消费账单ID
					ConsumeBillID: &bill.ID,       //消费所扣充值流水id
				}
				createConsumeBillsMap[bill.ID] = consumeBill
				//满足扣款条件
				var costValue int32
				updateBill := map[string]interface{}{
					"id": bill.ID,
				}
				if bill.GiftValueLeft >= consumeAmountLeft {
					//到当前流水够扣,否则就进入下一条数据
					costValue = consumeAmountLeft
					consumeAmountLeft = 0
					updateBill["gift_value_left"] = bill.GiftValueLeft - costValue //充值流水剩余赠金
					breakFlag = true
				} else {
					costValue = bill.GiftValueLeft //扣除金额等于当前水流剩余的赠金
					updateBill["gift_value_left"] = 0
					consumeAmountLeft -= costValue //更新剩余需要扣除的金额
				}
				consumeBill.GiftValue = costValue
				currRecord.GiftAmount += costValue
				currRecord.NewAccount.GiftValue -= costValue
				updateBillsMap[bill.ID] = updateBill
				if breakFlag {
					break
				}
			}
		}

		//赠金不足
		if consumeAmountLeft > 0 {
			//处理本金
			for _, bill := range usableValueBills {
				currRecord := billAccounts4Consume[*bill.AccountID] //肯定存在
				consumeBill, ok := createConsumeBillsMap[bill.ID]
				if !ok {
					//消费流水记录，可能还要更新baseValue，所以用map
					consumeBill = &model.TableConsumeBills{
						ID:            uuid.NewV4(),
						BillID:        &currRecord.ID, //消费账单ID
						ConsumeBillID: &bill.ID,       //消费所扣充值流水id
					}
					createConsumeBillsMap[bill.ID] = consumeBill
				}

				var costBaseValue int32 //当前循环抵扣的金额
				updateBill, ok := updateBillsMap[bill.ID]
				if !ok {
					updateBill = map[string]interface{}{
						"id": bill.ID,
					}
				}
				if bill.BaseValueLeft >= consumeAmountLeft {
					//到当前流水够扣,否则就进入下一条数据
					costBaseValue = consumeAmountLeft
					consumeAmountLeft = 0
					updateBill["base_value_left"] = bill.BaseValueLeft - costBaseValue //充值流水剩余本金
					breakFlag = true
				} else {
					costBaseValue = bill.BaseValueLeft //扣除金额等于当前流水剩余的本金
					updateBill["base_value_left"] = 0
					consumeAmountLeft -= costBaseValue //更新剩余需要扣除的金额
				}

				consumeBill.BaseValue = costBaseValue
				currRecord.BaseAmount += costBaseValue
				currRecord.NewAccount.BaseValue -= costBaseValue
				if breakFlag {
					break
				}
			}
		}

		if consumeAmountLeft > 0 {
			//余额不足
			iErr.ErrorMessage = "余额不足"
			iErr.ErrorCode = pkgs.ErrUnprocessableEntity
			return nil, iErr
		}
	}

	//商品
	if len(*request.Products) > 0 {
		var productIDs []uuid.UUID //需要消费的商品IDs
		for _, p := range *request.Products {
			productIDs = append(productIDs, p.ID)
		}

		//可用bill
		var usableProductBills []model.TableProductPackage
		subQuery := request.Tx.Model(&model.TableAccountBill{}).
			Select("id").
			Where("account_id in (?) AND change_category = 'recharge' AND after_account::json->>'products' is not null", request.AccountIDs)
		err := request.Tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_package_id in (?)", productIDs).
			Where("bill_id in (?) AND \"left\" > 0", subQuery).
			Order("created_at asc").
			Preload("AccountBill").
			//Preload("AccountBill.Account").
			Find(&usableProductBills).Error
		if err != nil {
			util.Logger.Error(fmt.Sprintf("consume 查询商品可用流水数据库错误:%+v", err))
			iErr.ErrorMessage = "数据库错误"
			iErr.ErrorCode = pkgs.ErrInternal
			return nil, iErr
		}
		//按照商品ID进行分类
		var productBillsMap map[uuid.UUID][]model.TableProductPackage //key是 product的ID
		for _, bill := range usableProductBills {
			if _, ok := productBillsMap[bill.ProductPackageID]; !ok {
				productBillsMap[bill.ProductPackageID] = []model.TableProductPackage{}
			}
			productBillsMap[bill.ProductPackageID] = append(productBillsMap[bill.ProductPackageID], bill)
		}

		for _, product := range *request.Products {
			if _, ok := productBillsMap[product.ID]; !ok {
				//不存在，就是余额不足
				util.Logger.Error(fmt.Sprintf("consume 商品余额不足:%s", product.ID))
				iErr.ErrorMessage = "商品库存不足"
				iErr.ErrorCode = pkgs.ErrInternal
				return nil, iErr
			}
			//开始扣商品
			productBalance := product.Count
			breakFlag := false

			for _, productBill := range productBillsMap[product.ID] {
				//account表数据需要更新
				var productCostCount int32
				if _, ok := billAccounts4Consume[*productBill.AccountBill.AccountID]; !ok {
					accountInfo := dealAccounts[*productBill.AccountBill.AccountID]
					if accountInfo.Status != model.AccountStatusActivated {
						util.Logger.Error(fmt.Sprintf("账户状态异常:%+v", accountInfo))
						iErr.ErrorMessage = "账户状态异常"
						iErr.ErrorCode = pkgs.ErrUnprocessableEntity
						return nil, iErr
					}
					newAccount := new(model.TableCardAccount)
					*newAccount = accountInfo
					billAccounts4Consume[*productBill.AccountBill.AccountID] = &billConsumeRecord{
						ID:         uuid.NewV4(),
						OldAccount: &accountInfo,
						NewAccount: newAccount,
					}
				}

				productBalance -= productBill.Left
				productCostCount = productBill.Left
				if productBalance <= 0 {
					//够扣了，结束循环
					productCostCount = product.Count
					productUpdateBill := make(map[string]interface{})
					productUpdateBill["id"] = productBill.ID
					productUpdateBill["left"] = productBill.Left - productBalance
					productUpdateBills = append(productUpdateBills, productUpdateBill)
					breakFlag = true
				} else {
					//本记录扣完
					productNoLeftBillIDs = append(productNoLeftBillIDs, productBill.ID)
				}
				//处理account表的products中对应的product
				updateAccount := billAccounts4Consume[*productBill.AccountBill.AccountID].NewAccount
				currProduct := billAccounts4Consume[*productBill.AccountBill.AccountID].Product
				*currProduct = append(*currProduct, fields.ProductPackageTicket{ID: product.ID, Count: productCostCount})
				for i, pr := range *updateAccount.Products {
					if pr.ID == productBill.ProductPackageID {
						// 减, afterAccount的products处理
						(*updateAccount.Products)[i].Count -= productCostCount
						break
					}
				}

				if breakFlag {
					break
				}
			}

			if productBalance > 0 {
				//全部扣完了还是不够，库存不足
				util.Logger.Error(fmt.Sprintf("consume 商品库存不足，还差:%d", productBalance))
				iErr.ErrorMessage = "商品库存不足"
				iErr.ErrorCode = pkgs.ErrInternal
				return nil, iErr
			}
		}
	}

	//套餐
	if len(*request.Packages) > 0 {
		var packageIDs []uuid.UUID //需要消费的商品IDs
		for _, p := range *request.Packages {
			packageIDs = append(packageIDs, p.ID)
		}

		//可用bill
		var usablePackageBills []model.TableProductPackage
		subQuery := request.Tx.Model(&model.TableAccountBill{}).
			Select("id").
			Where("account_id in (?) AND change_category = 'recharge' AND after_account::json->>'packages' is not null", request.AccountIDs)
		err := request.Tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_package_id in (?)", packageIDs).
			Where("bill_id in (?) AND left > 0", subQuery).
			Order("created_at asc").
			Preload("AccountBill").
			//Preload("AccountBill.Account").
			Find(&usablePackageBills).Error
		if err != nil {
			util.Logger.Error(fmt.Sprintf("consume 查询套餐可用流水数据库错误:%+v", err))
			iErr.ErrorMessage = "数据库错误"
			iErr.ErrorCode = pkgs.ErrInternal
			return nil, iErr
		}
		//按照商品ID进行分类
		var packageBillsMap map[uuid.UUID][]model.TableProductPackage //key是 package的ID
		for _, bill := range usablePackageBills {
			if _, ok := packageBillsMap[bill.ProductPackageID]; !ok {
				packageBillsMap[bill.ProductPackageID] = []model.TableProductPackage{}
			}
			packageBillsMap[bill.ProductPackageID] = append(packageBillsMap[bill.ProductPackageID], bill)
		}

		for _, pkg := range *request.Packages {
			if _, ok := packageBillsMap[pkg.ID]; !ok {
				//不存在，就是余额不足
				util.Logger.Error(fmt.Sprintf("consume 套餐库存不足:%s", pkg.ID))
				iErr.ErrorMessage = "套餐库存不足"
				iErr.ErrorCode = pkgs.ErrInternal
				return nil, iErr
			}
			//开始扣商品
			packageBalance := pkg.Count
			breakFlag := false

			for _, packageBill := range packageBillsMap[pkg.ID] {
				//account表数据需要更新
				var packageCostCount int32
				if _, ok := billAccounts4Consume[*packageBill.AccountBill.AccountID]; !ok {
					accountInfo := dealAccounts[*packageBill.AccountBill.AccountID]
					if accountInfo.Status != model.AccountStatusActivated {
						util.Logger.Error(fmt.Sprintf("账户状态异常:%+v", accountInfo))
						iErr.ErrorMessage = "账户状态异常"
						iErr.ErrorCode = pkgs.ErrUnprocessableEntity
						return nil, iErr
					}
					newAccount := new(model.TableCardAccount)
					*newAccount = accountInfo
					billAccounts4Consume[*packageBill.AccountBill.AccountID] = &billConsumeRecord{
						ID:         uuid.NewV4(),
						OldAccount: &accountInfo,
						NewAccount: newAccount,
					}
				}

				packageBalance -= packageBill.Left
				packageCostCount = packageBill.Left
				if packageBalance <= 0 {
					//够扣了，结束循环
					packageCostCount = pkg.Count
					packageUpdateBill := make(map[string]interface{})
					packageUpdateBill["id"] = packageBill.ID
					packageUpdateBill["left"] = packageBill.Left - packageBalance
					productUpdateBills = append(productUpdateBills, packageUpdateBill)
					breakFlag = true
				} else {
					//本记录扣完
					productNoLeftBillIDs = append(productNoLeftBillIDs, packageBill.ID)
				}
				//处理account表的products中对应的product
				updateAccount := billAccounts4Consume[*packageBill.AccountBill.AccountID].NewAccount
				currProduct := billAccounts4Consume[*packageBill.AccountBill.AccountID].Product
				*currProduct = append(*currProduct, fields.ProductPackageTicket{ID: pkg.ID, Count: packageCostCount})
				for i, pr := range *updateAccount.Packages {
					if pr.ID == packageBill.ProductPackageID {
						// 减, afterAccount的packages处理
						(*updateAccount.Packages)[i].Count -= packageCostCount
						break
					}
				}

				if breakFlag {
					break
				}
			}

			if packageBalance > 0 {
				//全部扣完了还是不够，库存不足
				util.Logger.Error(fmt.Sprintf("consume 商品库存不足，还差:%d", packageBalance))
				iErr.ErrorMessage = "商品库存不足"
				iErr.ErrorCode = pkgs.ErrInternal
				return nil, iErr
			}
		}
	}

	//一次消费 billCode相同
	billCode, err := s.getBillCode(ctx, request.BranchID, time.Now())
	if err != nil {
		iErr.ErrorCode = pkgs.ErrInternal
		iErr.ErrorMessage = "billCode 生成失败"
		return nil, iErr
	}
	//流水表消费记录 数量等于涉及的account数
	for accountID, record := range billAccounts4Consume {
		currentAccountID := accountID
		updateAccount := map[string]interface{}{
			"id":         currentAccountID,
			"base_value": gorm.Expr("base_value - ?", record.BaseAmount),
			"gift_value": gorm.Expr("gift_value - ?", record.GiftAmount),
		}
		if record.Product != nil {
			for _, product := range *record.Product {
				updateAccount[fmt.Sprintf("products::json->%s", product.ID)] = gorm.Expr(fmt.Sprintf("products::json->%s - ?", product.ID), product.Count)
			}
		}
		if record.Package != nil {
			for _, pkg := range *record.Package {
				updateAccount[fmt.Sprintf("packages::json->%s", pkg.ID)] = gorm.Expr(fmt.Sprintf("packages::json->%s - ?", pkg.ID), pkg.Count)
			}
		}

		billTable := &model.TableAccountBill{
			ID:             record.ID,
			AccountID:      &currentAccountID,
			BillCode:       billCode,
			CardID:         &cardInfo.ID,
			CardCode:       cardInfo.Code,
			BranchID:       &request.BranchID,
			ChangeValue:    record.GiftAmount + record.BaseAmount,
			ChangeCategory: model.BillCategoryConsume,
			ChangeType:     model.BillType(request.BillType),
			GiftValue:      record.GiftAmount,
			BaseValue:      record.BaseAmount,
			AfterAccount:   record.NewAccount,
			StaffID:        &request.StaffID,
			MerchantID:     &request.MerchantID,
		}
		if request.PosBillID != uuid.Nil {
			billTable.PosBillID = &request.PosBillID
		}
		createBills4Consume = append(createBills4Consume, billTable)
		updateAccounts = append(updateAccounts, updateAccount)
	}
	err = model.Consume(request.Tx, updateBillsMap, createBills4Consume, createConsumeBillsMap, productNoLeftBillIDs, productUpdateBills, updateAccounts)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("Consume 数据库错误:%v", err))
		iErr.ErrorMessage = "数据库错误"
		iErr.ErrorCode = pkgs.ErrInternal
		return nil, iErr
	}

	return createConsumeBillsMap, iErr
}

//获取转账入账账户，可能会没有
func getTransferDestAccount(sourceAccount *model.TableCardAccount, cardID, merchantID uuid.UUID) (*model.TableCard, *model.TableCardAccount, internalErr) {
	var iErr internalErr
	card, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			iErr.ErrorMessage = "卡不存在"
			iErr.ErrorCode = pkgs.ErrNotFound
			return nil, nil, iErr
		}
		util.Logger.Error(fmt.Sprintf("RechargeCard 查询卡数据库错误:%v", err))
		iErr.ErrorMessage = "数据库错误"
		iErr.ErrorCode = pkgs.ErrInternal
		return nil, nil, iErr
	}

	if card.Status != model.CardStatusActive {
		util.Logger.Error(fmt.Sprintf("transfer 到账卡状态错误:%s", card.Status))
		iErr.ErrorMessage = "卡状态错误"
		iErr.ErrorCode = pkgs.ErrUnprocessableEntity
		return nil, nil, iErr
	}
	accounts, err := model.GetCardAccountsByCardID(card.ID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("transfer 查询账户数据库错误:%v", err))
		iErr.ErrorMessage = "数据库错误"
		iErr.ErrorCode = pkgs.ErrInternal
		return nil, nil, iErr
	}

	for _, v := range accounts {
		if uuid.Equal(*v.TagID, *sourceAccount.TagID) && v.Status == model.AccountStatusActivated {
			if v.ID == sourceAccount.ID {
				iErr.ErrorMessage = "不能转账到原账户"
				iErr.ErrorCode = pkgs.ErrUnprocessableEntity
				return nil, nil, iErr
			}
			return card, &v, iErr
		}
	}

	return card, nil, iErr
}
