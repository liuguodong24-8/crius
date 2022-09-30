package rpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	basicProto "gitlab.omytech.com.cn/micro-service/basic/proto"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
	"gitlab.omytech.com.cn/micro-service/member-account/util"
	memberExtension "gitlab.omytech.com.cn/micro-service/member-extension/proto"
	private "gitlab.omytech.com.cn/micro-service/member-private/proto"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

type internalErr struct {
	ErrorCode    int32
	ErrorMessage string
}

// MakeCard 制卡
func (s *Server) MakeCard(ctx context.Context, request *proto.MakeCardRequest) (*proto.MakeCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("MakeCard")
	resp := &proto.MakeCardResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)
	operatorID := pkgs.GetMetadata(ctx).StaffID
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	if branchID == uuid.Nil || operatorID == uuid.Nil || merchantID == uuid.Nil ||
		(request.Category != "member" && request.Category != "gift") || request.Code == "" {
		cutil.Logger.Error("MakeCard rpc请求参数错误")
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	_, err := model.ShowCardByCode(request.Code, merchantID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			cutil.Logger.Error(fmt.Sprintf("MakeCard 查询卡数据库错误:%v", err))
			resp.ErrorMessage = "制卡请求失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	} else {
		resp.ErrorMessage = "卡号已存在"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	card := model.TableCard{
		ID:             uuid.NewV4(),
		Category:       model.CardCategory(request.Category),
		Code:           request.Code,
		CreateBranchID: &branchID,
		CreateStaffID:  &operatorID,
		Status:         model.CardStatusInit,
		MerchantID:     &merchantID,
	}

	err = model.CreateCard(card)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("MakeCard 新增卡数据库错误:%v", err))
		resp.ErrorMessage = "制卡请求失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	after := pkgs.MakeParams(card)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: card.TableName(),
		After:             &after,
		TableID:           &card.ID,
		Method:            "create",
	}

	return resp, nil
}

// ActivePrimaryCard 开主卡
func (s *Server) ActivePrimaryCard(ctx context.Context, request *proto.ActivePrimaryCardRequest) (*proto.ActivePrimaryCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ActivePrimaryCard")
	resp := &proto.ActivePrimaryCardResponse{
		ErrorCode: pkgs.Success,
	}

	operatorID := pkgs.GetMetadata(ctx).StaffID
	branchID := uuid.FromStringOrNil(request.BranchId)
	cardID := uuid.FromStringOrNil(request.CardId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	// 查询卡状态
	var tabCard *model.TableCard
	tabCard, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			cutil.Logger.Error(fmt.Sprintf("ActivePrimaryCard 卡不存在:%v", cardID))
			return resp, nil
		}
		cutil.Logger.Error(fmt.Sprintf("ActivePrimaryCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "激活卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if tabCard.Status != model.CardStatusInit {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "卡已激活"
		return resp, nil
	}
	beforeCard := pkgs.MakeParams(*tabCard)

	money, promotionOptions, ppt, err := s.listPromotionOption(ctx, request.Promotions, request.RechargeValue, branchID)
	if err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	memberResp, err := s.merchantBasic().CreateMember(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.CreateMemberRequest{
		Member: &merchantBasic.MemberInfo{
			Name:          request.Name,
			Phone:         request.Phone,
			PhoneCode:     request.PhoneCode,
			Gender:        request.Gender,
			FirstBranchId: branchID.String(),
			Channel:       "open_card",
			Birthday:      request.Birthday,
		},
	})
	if err != nil || memberResp == nil || memberResp.ErrorCode != pkgs.Success {
		cutil.Logger.Error(fmt.Sprintf("ActivePrimaryCard 新增会员错误:%v, %v", err, memberResp))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	cardAccount, err := activePrimaryCardData(uuid.FromStringOrNil(memberResp.Data), branchID, merchantID, operatorID, tabCard, util.Sha256(request.Password), money)
	if err != nil {
		if err == model.ErrAccountAbnormal {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "账户状态异常"
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}
	now := time.Now()

	// 生成账单号
	billCode, err := s.getBillCode(ctx, branchID, now)
	if err != nil || billCode == "" {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	payments := model.Payments{
		Wechat: request.Payments.Wechat,
		Cash:   request.Payments.Cash,
		Card:   request.Payments.Card,
		Alipay: request.Payments.Alipay,
	}
	afterAccount := pkgs.MakeParams(cardAccount)
	bill := model.TableAccountBill{
		ID:               uuid.NewV4(),
		BillCode:         billCode,
		AccountID:        &cardAccount.ID,
		CardID:           &tabCard.ID,
		CardCode:         tabCard.Code,
		BranchID:         &branchID,
		ChangeValue:      request.RechargeValue,
		ChangeCategory:   model.BillCategoryRecharge,
		ChangeType:       model.BillTypeOpen,
		BaseValue:        money.baseValue,
		GiftValue:        money.giftValue,
		Payments:         &payments,
		StaffID:          &operatorID,
		OperatorComment:  request.Recommender,
		MerchantID:       &merchantID,
		BaseValueLeft:    money.baseValue,
		GiftValueLeft:    money.giftValue,
		PromotionOptions: &promotionOptions,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}

	account := *cardAccount
	//CreatedAt 则为更新账户，修改本金赠金为充值金额
	if account.CreatedAt != nil {
		account.BaseValue, account.GiftValue = money.baseValue, money.giftValue
	}

	for _, v := range ppt.tickets {
		params := make([]*memberExtension.MemberCoupon, v.Number)
		for i := int32(0); i < v.Number; i++ {
			params[i] = &memberExtension.MemberCoupon{
				CouponId: v.Id,
				MemberId: account.MemberID.String(),
			}
		}
		couponResp, err := s.memberExtension().CreateMemberCoupons(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &memberExtension.CreateMemberCouponsRequest{
			MemberCoupons: params,
		})
		if err != nil || couponResp.ErrorCode != pkgs.Success {
			cutil.Logger.Error(fmt.Sprintf("ActivePrimaryCard 更新卡激活状态错误:%v", err))
			resp.ErrorMessage = "激活卡失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	}

	var products, packages []model.TableProductPackage
	for _, v := range ppt.products {
		products = append(products, model.TableProductPackage{
			ID:               uuid.NewV4(),
			ProductPackageID: uuid.FromStringOrNil(v.Id),
			Code:             v.Code,
			Number:           v.Number,
			Price:            v.Price,
			Title:            v.Title,
			BillID:           &bill.ID,
			Left:             v.Number,
			Category:         model.ProductPackageCategory(v.Category),
		})
	}
	for _, v := range ppt.packages {
		packages = append(packages, model.TableProductPackage{
			ID:               uuid.NewV4(),
			ProductPackageID: uuid.FromStringOrNil(v.Id),
			Code:             v.Code,
			Number:           v.Number,
			Price:            v.Price,
			Title:            v.Title,
			BillID:           &bill.ID,
			Left:             v.Number,
			Category:         model.ProductPackageCategory(v.Category),
		})
	}
	accountProductPackage(&account, ppt)
	bill.AfterAccount = &account

	if err = model.RechargeCard(*tabCard, account, bill, products, packages); err != nil {
		cutil.Logger.Error(fmt.Sprintf("ActivePrimaryCard 更新卡激活状态错误:%v", err))
		resp.ErrorMessage = "激活卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	go model.SaveBillPromotion(bill)

	staffID := pkgs.GetMetadata(ctx).StaffID
	afterCard := pkgs.MakeParams(tabCard)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: tabCard.TableName(),
		Before:            &beforeCard,
		After:             &afterCard,
		TableID:           &tabCard.ID,
		Method:            "update",
	}
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: cardAccount.TableName(),
		After:             &afterAccount,
		TableID:           &cardAccount.ID,
		Method:            "update",
	}
	resp.Data = toProtoBillData(bill)
	resp.Data.Products = ppt.products
	resp.Data.Packages = ppt.packages
	resp.Data.TotalProducts = toAccountProto(&account).Products
	resp.Data.TotalPackages = toAccountProto(&account).Packages

	return resp, nil
}

// ActiveSecondaryCard 开副卡
func (s *Server) ActiveSecondaryCard(ctx context.Context, request *proto.ActiveSecondaryCardRequest) (*proto.ActiveSecondaryCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ActiveSecondaryCard")
	resp := &proto.ActiveSecondaryCardResponse{
		ErrorCode: pkgs.Success,
	}

	operatorID := pkgs.GetMetadata(ctx).StaffID
	branchID := uuid.FromStringOrNil(request.BranchId)
	cardID := uuid.FromStringOrNil(request.CardId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	primaryID := uuid.FromStringOrNil(request.PrimaryId)
	primaryAccountID := uuid.FromStringOrNil(request.PrimaryAccountId)

	tabCard, primaryCard, resp := activeSecondaryCardValidate(cardID, primaryID, primaryAccountID, merchantID, util.Sha256(request.PrimaryPassword), request.PrimaryVerified)
	if resp.ErrorCode != pkgs.Success {
		return resp, nil
	}
	beforeCard := pkgs.MakeParams(*tabCard)
	tabCard.PrimaryID = &primaryID

	primaryAccount, err := model.ShowCardAccount(primaryAccountID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "主卡账户不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return resp, nil
		}
		cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 查询账户数据库错误:%v", err))
		resp.ErrorMessage = "开副卡失败"
		resp.ErrorCode = pkgs.ErrPassword
		return resp, nil
	}

	if primaryAccount.Status != model.AccountStatusActivated {
		resp.ErrorMessage = "主卡账户异常"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	cardAccount := model.TableCardAccount{
		ID:         uuid.NewV4(),
		BranchID:   &branchID,
		Status:     model.AccountStatusActivated,
		MerchantID: merchantID,
		TagID:      primaryAccount.TagID,
		Category:   model.CardSubCategorySecondary,
	}
	now := time.Now()
	tabCard.Category = model.CardCategoryMember
	tabCard.SubCategory = model.CardSubCategorySecondary
	tabCard.OpenBranchID = &branchID
	tabCard.AccountIDs = &fields.UUIDArr{cardAccount.ID}
	tabCard.Status = model.CardStatusActive
	tabCard.OpenedAt = &now
	tabCard.OpenStaffID = &operatorID
	if request.Password == "" {
		tabCard.Password = util.DefaultCardPassword
	} else {
		tabCard.Password = util.Sha256(request.Password)
	}

	// 生成账单号
	primaryCode, err := s.getBillCode(ctx, branchID, now)
	if err != nil || primaryCode == "" {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}
	secondaryCode, err := s.getBillCode(ctx, branchID, now)
	if err != nil || secondaryCode == "" {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	primaryBill := model.TableAccountBill{
		ID:              uuid.NewV4(),
		BillCode:        primaryCode,
		AccountID:       &primaryAccountID,
		CardID:          &primaryCard.ID,
		CardCode:        primaryCard.Code,
		BranchID:        &branchID,
		ChangeValue:     request.RechargeValue,
		ChangeCategory:  model.BillCategoryConsume,
		ChangeType:      model.BillTypeTransfer,
		StaffID:         &operatorID,
		OperatorComment: request.Recommender,
		MerchantID:      &merchantID,
		CreatedAt:       &now,
		UpdatedAt:       &now,
	}

	secondaryBill := model.TableAccountBill{
		ID:              uuid.NewV4(),
		BillCode:        secondaryCode,
		AccountID:       &cardAccount.ID,
		CardID:          &tabCard.ID,
		CardCode:        tabCard.Code,
		BranchID:        &branchID,
		ChangeValue:     request.RechargeValue,
		ChangeCategory:  model.BillCategoryRecharge,
		ChangeType:      model.BillTypeSub,
		StaffID:         &operatorID,
		OperatorComment: request.Recommender,
		MerchantID:      &merchantID,
		CreatedAt:       &now,
		UpdatedAt:       &now,
	}

	if err := model.ActiveSecondaryCardTransaction(tabCard, primaryAccount, &cardAccount, &primaryBill, &secondaryBill, primaryAccountID, request.RechargeValue); err != nil {
		if err == model.ErrBeyondPrimaryAccountBalance {
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
			resp.ErrorMessage = "副卡金额不可大于账户可用余额"
			return resp, nil
		}
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	afterCard := pkgs.MakeParams(tabCard)
	afterAccount := pkgs.MakeParams(cardAccount)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: tabCard.TableName(),
		Before:            &beforeCard,
		After:             &afterCard,
		TableID:           &tabCard.ID,
		Method:            "update",
	}
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: cardAccount.TableName(),
		Before:            nil,
		After:             &afterAccount,
		TableID:           &cardAccount.ID,
		Method:            "update",
	}
	resp.Data = toProtoBillData(secondaryBill)
	resp.Data.PrimaryId = primaryCard.ID.String()
	resp.Data.PrimaryCode = primaryCard.Code
	return resp, nil
}

// ActiveBlankCard 开不记名卡
func (s *Server) ActiveBlankCard(ctx context.Context, request *proto.ActiveBlankCardRequest) (*proto.ActiveBlankCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ActiveBlankCard")
	resp := &proto.ActiveBlankCardResponse{
		ErrorCode: pkgs.Success,
	}

	operatorID := pkgs.GetMetadata(ctx).StaffID
	branchID := uuid.FromStringOrNil(request.BranchId)
	cardID := uuid.FromStringOrNil(request.CardId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	tabCard, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			cutil.Logger.Error(fmt.Sprintf("ActiveblankCard 卡不存在:%v", cardID))
			return resp, nil
		}
		cutil.Logger.Error(fmt.Sprintf("ActiveblankCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "激活卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	beforeCard := pkgs.MakeParams(*tabCard)
	if tabCard.Status != model.CardStatusInit {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "卡已初始化"
		return resp, nil
	}

	money, promotionOptions, ppt, err := s.listPromotionOption(ctx, request.Promotions, request.RechargeValue, branchID)
	if err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	cardAccount := model.TableCardAccount{
		ID:         uuid.NewV4(),
		BaseValue:  money.baseValue,
		GiftValue:  money.giftValue,
		BranchID:   &branchID,
		Status:     model.AccountStatusActivated,
		MerchantID: merchantID,
		TagID:      &money.tagID,
		Category:   model.CardSubCategoryBlank,
	}
	now := time.Now()
	tabCard.Category = model.CardCategoryMember
	tabCard.SubCategory = model.CardSubCategoryBlank
	tabCard.OpenBranchID = &branchID
	tabCard.AccountIDs = &fields.UUIDArr{cardAccount.ID}
	tabCard.Status = model.CardStatusActive
	tabCard.OpenedAt = &now
	tabCard.OpenStaffID = &operatorID
	if request.Password == "" {
		tabCard.Password = util.DefaultCardPassword
	} else {
		tabCard.Password = util.Sha256(request.Password)
	}

	// 生成账单号
	billCode, err := s.getBillCode(ctx, branchID, now)
	if err != nil || billCode == "" {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "激活卡失败"
		return resp, nil
	}

	payments := model.Payments{
		Wechat: request.Payments.Wechat,
		Cash:   request.Payments.Cash,
		Card:   request.Payments.Card,
		Alipay: request.Payments.Alipay,
	}
	afterAccount := pkgs.MakeParams(cardAccount)
	bill := model.TableAccountBill{
		ID:               uuid.NewV4(),
		BillCode:         billCode,
		AccountID:        &cardAccount.ID,
		CardID:           &tabCard.ID,
		CardCode:         tabCard.Code,
		BranchID:         &branchID,
		ChangeValue:      request.RechargeValue,
		ChangeCategory:   model.BillCategoryRecharge,
		ChangeType:       model.BillTypeNobody,
		BaseValue:        money.baseValue,
		GiftValue:        money.giftValue,
		Payments:         &payments,
		StaffID:          &operatorID,
		OperatorComment:  request.Recommender,
		MerchantID:       &merchantID,
		BaseValueLeft:    money.baseValue,
		GiftValueLeft:    money.giftValue,
		PromotionOptions: &promotionOptions,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}

	var products, packages []model.TableProductPackage
	for _, v := range ppt.products {
		products = append(products, model.TableProductPackage{
			ID:               uuid.NewV4(),
			ProductPackageID: uuid.FromStringOrNil(v.Id),
			Code:             v.Code,
			Number:           v.Number,
			Price:            v.Price,
			Title:            v.Title,
			BillID:           &bill.ID,
			Left:             v.Number,
			Category:         model.ProductPackageCategory(v.Category),
		})
	}
	for _, v := range ppt.packages {
		packages = append(packages, model.TableProductPackage{
			ID:               uuid.NewV4(),
			ProductPackageID: uuid.FromStringOrNil(v.Id),
			Code:             v.Code,
			Number:           v.Number,
			Price:            v.Price,
			Title:            v.Title,
			BillID:           &bill.ID,
			Left:             v.Number,
			Category:         model.ProductPackageCategory(v.Category),
		})
	}
	accountProductPackage(&cardAccount, ppt)
	bill.AfterAccount = &cardAccount
	err = model.RechargeCard(*tabCard, cardAccount, bill, products, packages)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ActiveblankCard 更新卡激活状态错误:%v", err))
		resp.ErrorMessage = "激活卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	go model.SaveBillPromotion(bill)

	staffID := pkgs.GetMetadata(ctx).StaffID
	afterCard := pkgs.MakeParams(tabCard)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: tabCard.TableName(),
		Before:            &beforeCard,
		After:             &afterCard,
		TableID:           &tabCard.ID,
		Method:            "update",
	}
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: cardAccount.TableName(),
		Before:            nil,
		After:             &afterAccount,
		TableID:           &cardAccount.ID,
		Method:            "update",
	}
	resp.Data = toProtoBillData(bill)
	resp.Data.Products = ppt.products
	resp.Data.Packages = ppt.packages
	resp.Data.TotalProducts = toAccountProto(&cardAccount).Products
	resp.Data.TotalPackages = toAccountProto(&cardAccount).Packages
	return resp, nil
}

// RechargeCard 充值
func (s *Server) RechargeCard(ctx context.Context, request *proto.RechargeCardRequest) (*proto.RechargeCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("RechargeCard")
	resp := &proto.RechargeCardResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(request.BranchId)
	operatorID := pkgs.GetMetadata(ctx).StaffID
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	cardID := uuid.FromStringOrNil(request.CardId)

	card, resp := rechargeCardValidate(cardID, merchantID)
	if resp.ErrorCode != pkgs.Success {
		return resp, nil
	}

	money, promotionOptions, ppt, err := s.listPromotionOption(ctx, request.Promotions, request.RechargeValue, branchID)
	if err != nil {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "充值失败"
		return resp, nil
	}

	accounts, err := model.GetAccountByMemberID(*card.MemberID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("RechargeCard 查询账户数据库错误:%v", err))
		resp.ErrorMessage = "充值失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var memberID *uuid.UUID
	var cardAccount model.TableCardAccount
	for _, v := range accounts {
		memberID = v.MemberID
		if *v.TagID == money.tagID && v.Category == model.CardSubCategoryPrimary {
			if v.Status != model.AccountStatusActivated {
				resp.ErrorCode = pkgs.ErrUnprocessableEntity
				resp.ErrorMessage = "账户状态异常"
				return resp, nil
			}
			cardAccount = v
			break
		}
	}

	if cardAccount.ID == uuid.Nil {
		cardAccount = model.TableCardAccount{
			ID:         uuid.NewV4(),
			MemberID:   memberID,
			BranchID:   &branchID,
			Status:     model.AccountStatusActivated,
			MerchantID: merchantID,
			Category:   model.CardSubCategoryPrimary,
			TagID:      &money.tagID,
		}
	}

	beforeAccount := pkgs.MakeParams(cardAccount)
	cardAccount.BaseValue += money.baseValue
	cardAccount.GiftValue += money.giftValue

	//TODO 赠送商品
	payments := model.Payments{
		Wechat: request.Payments.Wechat,
		Cash:   request.Payments.Cash,
		Card:   request.Payments.Card,
		Alipay: request.Payments.Alipay,
	}
	afterAccount := pkgs.MakeParams(cardAccount)

	// 生成账单号
	billCode, err := s.getBillCode(ctx, branchID, time.Now())
	if err != nil || billCode == "" {
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "充值失败"
		return resp, nil
	}

	now := time.Now()
	bill := model.TableAccountBill{
		ID:               uuid.NewV4(),
		BillCode:         billCode,
		AccountID:        &cardAccount.ID,
		BranchID:         &branchID,
		ChangeValue:      request.RechargeValue,
		ChangeCategory:   model.BillCategoryRecharge,
		ChangeType:       model.BillTypeRecharge,
		BaseValue:        money.baseValue,
		GiftValue:        money.giftValue,
		CardID:           &cardID,
		CardCode:         card.Code,
		Payments:         &payments,
		StaffID:          &operatorID,
		OperatorComment:  request.Recommender,
		MerchantID:       &merchantID,
		BaseValueLeft:    money.baseValue,
		GiftValueLeft:    money.giftValue,
		PromotionOptions: &promotionOptions,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}

	if !cutil.ArrContainElement(card.AccountIDs.Slice(), cardAccount.ID) {
		if card.AccountIDs == nil {
			card.AccountIDs = &fields.UUIDArr{}
		}
		*card.AccountIDs = append(*card.AccountIDs, cardAccount.ID)
	}

	account := cardAccount
	//CreatedAt 则为更新账户，修改本金赠金为充值金额
	if account.CreatedAt != nil {
		account.BaseValue, account.GiftValue = money.baseValue, money.giftValue
	}
	var products, packages []model.TableProductPackage
	for _, v := range ppt.products {
		products = append(products, model.TableProductPackage{
			ID:               uuid.NewV4(),
			ProductPackageID: uuid.FromStringOrNil(v.Id),
			Code:             v.Code,
			Number:           v.Number,
			Price:            v.Price,
			Title:            v.Title,
			BillID:           &bill.ID,
			Left:             v.Number,
			Category:         model.ProductPackageCategory(v.Category),
		})
	}
	for _, v := range ppt.packages {
		packages = append(packages, model.TableProductPackage{
			ID:               uuid.NewV4(),
			ProductPackageID: uuid.FromStringOrNil(v.Id),
			Code:             v.Code,
			Number:           v.Number,
			Price:            v.Price,
			Title:            v.Title,
			BillID:           &bill.ID,
			Left:             v.Number,
			Category:         model.ProductPackageCategory(v.Category),
		})
	}
	accountProductPackage(&account, ppt)
	bill.AfterAccount = &account
	err = model.RechargeCard(*card, account, bill, products, packages)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("RechargeCard 更新账户数据库错误:%v", err))
		resp.ErrorMessage = "充值失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	go model.SaveBillPromotion(bill)

	staffID := pkgs.GetMetadata(ctx).StaffID
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: cardAccount.TableName(),
		Before:            &beforeAccount,
		After:             &afterAccount,
		TableID:           &cardAccount.ID,
		Method:            "update",
	}
	resp.Data = toProtoBillData(bill)
	resp.Data.Products = ppt.products
	resp.Data.Packages = ppt.packages
	resp.Data.TotalProducts = toAccountProto(&account).Products
	resp.Data.TotalPackages = toAccountProto(&account).Packages
	return resp, nil
}

// BindCard 绑卡 此接口没有使用
func (s *Server) BindCard(ctx context.Context, request *proto.BindCardRequest) (*proto.BindCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("BindCard")
	resp := &proto.BindCardResponse{
		ErrorCode: pkgs.Success,
	}

	operatorID := pkgs.GetMetadata(ctx).StaffID
	branchID := uuid.FromStringOrNil(request.BranchId)
	accountID := uuid.FromStringOrNil(request.AccountId)
	cardID := uuid.FromStringOrNil(request.CardId)
	merchantID := pkgs.GetMetadata(ctx).MerchantID

	tabCard, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			cutil.Logger.Error(fmt.Sprintf("BindCard 卡不存在:%v", request.CardId))
			return resp, nil
		}
		cutil.Logger.Error(fmt.Sprintf("BindCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "绑卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if tabCard.Status != model.CardStatusInit {
		resp.ErrorMessage = "卡已激活"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	beforeCard := pkgs.MakeParams(*tabCard)

	account, err := model.ShowCardAccount(accountID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "账户不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			cutil.Logger.Error(fmt.Sprintf("BindCard 账户不存在:%v", request.AccountId))
			return resp, nil
		}
		cutil.Logger.Error(fmt.Sprintf("BindCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "绑卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if account.MerchantID != merchantID {
		resp.ErrorMessage = "账户不存在"
		resp.ErrorCode = pkgs.ErrNotFound
		cutil.Logger.Error("BindCard 账户商户id与上传商户id不一致")
		return resp, nil
	}

	if !cutil.ArrContainElement(tabCard.AccountIDs.Slice(), accountID) {
		*tabCard.AccountIDs = append(*tabCard.AccountIDs, accountID)
	}
	if request.Password == "" {
		tabCard.Password = util.DefaultCardPassword
	} else {
		tabCard.Password = util.Sha256(request.Password)
	}
	tabCard.Status = model.CardStatusActive
	now := time.Now()
	tabCard.OpenedAt = &now
	tabCard.OpenStaffID = &operatorID
	tabCard.OpenBranchID = &branchID
	tabCard.SubCategory = model.CardSubCategoryPrimary

	err = model.UpdateCard(*tabCard)
	if err != nil {
		cutil.Logger.Error("BindCard 更新卡数据库错误:%v")
		resp.ErrorMessage = "绑卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	staffID := pkgs.GetMetadata(ctx).StaffID
	afterCard := pkgs.MakeParams(tabCard)
	model.SnapShotChan <- &model.TableSnapshot{
		ID:                uuid.NewV4(),
		StaffID:           &staffID,
		SleuthCode:        pkgs.GetMetadata(ctx).SleuthCode,
		SnapShotTableName: tabCard.TableName(),
		Before:            &beforeCard,
		After:             &afterCard,
		TableID:           &tabCard.ID,
		Method:            "update",
	}
	return resp, nil
}

// GetCards 获取卡列表
func (s *Server) GetCards(ctx context.Context, request *proto.GetCardsRequest) (*proto.GetCardsResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetCards")
	resp := &proto.GetCardsResponse{
		ErrorCode: pkgs.Success,
	}
	//如果没有筛选门店，那就是账号所属的门店
	var branchIDs []uuid.UUID
	branchID := uuid.FromStringOrNil(request.BranchId)
	if branchID != uuid.Nil {
		branchIDs = append(branchIDs, branchID)
	} else {
		branchIDs = pkgs.GetMetadata(ctx).BranchIDs
	}

	count, err := model.CountCards(branchIDs, request.Category, request.Status)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("model.CountCards 数据库查询错误:%v", err))
		resp.ErrorMessage = "数据库查询错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if count == 0 {
		resp.Data = &proto.CardList{Total: 0}
		return resp, nil
	}
	cards, err := model.GetCards(branchIDs, request.Category, request.Status, request.Offset, request.Limit)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf(" model.GetCards 数据库查询错误:%v", err))
		resp.ErrorMessage = "数据库查询错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var cardData []*proto.Card
	for _, card := range cards {
		cardData = append(cardData, toCardProto(&card))
	}

	resp.Data = &proto.CardList{Cards: cardData, Total: int32(count)}
	return resp, nil
}

// ShowCard 根据code查询卡信息
func (s *Server) ShowCard(ctx context.Context, request *proto.ShowCardRequest) (*proto.ShowCardResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowCard")
	resp := &proto.ShowCardResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)
	if len(request.CardCode) == 0 {
		cutil.Logger.Error(fmt.Sprintf("ShowCard 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	card, err := model.ShowCardByCode(request.CardCode, metadata.MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("model.ShowCardByCode 数据库查询错误:%v", err))
		resp.ErrorMessage = "数据库查询错误"
		resp.ErrorCode = pkgs.ErrInternal
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "未查询到卡"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}

	resp.Data = toCardProto(card)
	return resp, nil
}

// ShowCardByID 根据id查询卡信息
func (s *Server) ShowCardByID(ctx context.Context, request *proto.ShowCardByIDRequest) (*proto.ShowCardByIDResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowCardByID")
	resp := &proto.ShowCardByIDResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)
	ID := uuid.FromStringOrNil(request.Id)
	if ID == uuid.Nil {
		cutil.Logger.Error(fmt.Sprintf("ShowCard 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	card, err := model.ShowCard(ID, metadata.MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("model.ShowCard 数据库查询错误:%v", err))
		resp.ErrorMessage = "数据库查询错误"
		resp.ErrorCode = pkgs.ErrInternal
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "未查询到卡"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}

	resp.Data = toCardProto(card)
	return resp, nil
}

// GetCardsByAccount 根据账户获取卡
func (s *Server) GetCardsByAccount(ctx context.Context, request *proto.GetCardsByAccountRequest) (*proto.GetCardsByAccountResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetCardsByAccount")
	resp := &proto.GetCardsByAccountResponse{
		ErrorCode: pkgs.Success,
	}
	accountID := uuid.FromStringOrNil(request.AccountId)
	if accountID == uuid.Nil {
		cutil.Logger.Error(fmt.Sprintf("GetCardsByAccount 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	cards, err := model.GetCardsByAccountID(accountID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("GetCardsByAccountID 数据库查询错误:%v", err))
		resp.ErrorMessage = "数据库查询错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var cardsProto []*proto.Card
	if len(cards) > 0 {
		for _, card := range cards {
			cardsProto = append(cardsProto, toCardProto(card))
		}
	}

	resp.Data = cardsProto
	return resp, nil
}

// UpdateCardStatus 卡挂失、找回
func (s *Server) UpdateCardStatus(ctx context.Context, request *proto.UpdateCardStatusRequest) (*proto.NoDataResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("UpdateCardStatus")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)
	cardID := uuid.FromStringOrNil(request.CardId)
	if request.Action == "" || cardID == uuid.Nil {
		cutil.Logger.Error(fmt.Sprintf("UpdateCardStatus 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	//先确认卡状态
	cardInfo, err := model.ShowCard(cardID, metadata.MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("UpdateCardStatus 查询数据库错误:%v", err))
		resp.ErrorMessage = "查询数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "对应的卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}

	nextStatus, err := validCardStatus(cardInfo.Status, request.Action)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("validCardStatus 卡状态异常, curr:%s, action:%s", cardInfo.Status, request.Action))
		resp.ErrorMessage = "卡状态异常"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	cardLost := model.TableCardLost{
		ID:         uuid.NewV4(),
		CardID:     &cardID,
		StaffID:    &metadata.StaffID,
		Action:     request.Action,
		MerchantID: &metadata.MerchantID,
	}
	err = model.UpdateCardStatus(&cardLost, nextStatus)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("更新卡状态失败 %v", err))
		resp.ErrorMessage = "更新卡状态失败"
		resp.ErrorCode = pkgs.ErrInternal
	}

	return resp, nil
}

// ReplaceCard 补卡
func (s *Server) ReplaceCard(ctx context.Context, request *proto.ReplaceCardRequest) (*proto.NoDataResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ReplaceCard")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)

	currCardID := uuid.FromStringOrNil(request.CurrCardId)
	newCardID := uuid.FromStringOrNil(request.NewCardId)
	branchID := uuid.FromStringOrNil(request.OpenBranchId)
	staffID := metadata.StaffID
	merchantID := metadata.MerchantID

	//if currCardID == uuid.Nil ||
	//	newCardID == uuid.Nil ||
	//	branchID == uuid.Nil ||
	//	request.ReplacementCost < 0 ||
	//	request.PayMethod == "" {
	//	cutil.Logger.Error(fmt.Sprintf("ReplaceCard 请求参数错误:%v", request))
	//	resp.ErrorMessage = "请求参数错误"
	//	resp.ErrorCode = pkgs.ErrUnprocessableEntity
	//	return resp, nil
	//}
	currCard, err := model.ShowCard(currCardID, metadata.MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("换卡查询旧卡信息数据库错误%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorMessage = "参数错误，旧卡不存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
		}
		return resp, nil
	}

	if currCard.Status != model.CardStatusActive {
		cutil.Logger.Error(fmt.Sprintf("旧卡状态异常:%s", currCard.Status))
		resp.ErrorMessage = "参数错误，旧卡非激活状态"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}
	updateCurrCard := model.TableCard{
		ID:     currCardID,
		Status: model.CardStatusLost,
	}

	newCard, err := model.ShowCard(newCardID, metadata.MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("换卡查询新卡信息数据库错误%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorMessage = "参数错误，新卡不存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
		}
		return resp, nil
	}
	if newCard.Status != model.CardStatusInit {
		cutil.Logger.Error(fmt.Sprintf("新卡状态异常:%s", newCard.Status))
		resp.ErrorMessage = "参数错误，新卡非初始化状态"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	password := util.DefaultCardPassword
	if len(request.Password) > 0 {
		password = util.Sha256(request.Password)
	}
	now := time.Now()
	updateNewCard := model.TableCard{
		ID:           newCardID,
		Category:     currCard.Category,
		SubCategory:  currCard.SubCategory,
		PrimaryID:    currCard.PrimaryID,
		OpenBranchID: &branchID,
		AccountIDs:   currCard.AccountIDs,
		MerchantID:   &metadata.MerchantID,
		Status:       model.CardStatusActive,
		OpenedAt:     &now,
		Password:     password,
		OpenStaffID:  &metadata.StaffID,
		MemberID:     currCard.MemberID,
	}
	tx := model.GetDBEntity().Conn.Begin()
	//卡余额支付
	if request.Payments.Card > 0 {
		consumeRequest := &consumeRequest{
			Tx:         tx,
			CardID:     currCardID,
			AccountIDs: currCard.AccountIDs.Slice(),
			CostValue:  request.Payments.Card,
			BranchID:   branchID,
			StaffID:    staffID,
			MerchantID: merchantID,
			BillType:   string(model.BillTypeReplace),
		}
		_, iErr := s.consume(ctx, consumeRequest)
		if iErr.ErrorCode != 0 {
			resp.ErrorCode = iErr.ErrorCode
			resp.ErrorMessage = iErr.ErrorMessage
			return resp, nil
		}

	}
	payments := pkgs.MakeParams(request.Payments)
	cardReplaceInsert := model.TableCardReplace{
		ID:         uuid.NewV4(),
		CurrCardID: &currCardID,
		NewCardID:  &newCardID,
		StaffID:    &metadata.StaffID,
		Payments:   &payments,
		MerchantID: &metadata.MerchantID,
	}
	err = model.ReplaceCard(tx, &updateCurrCard, &updateNewCard, &cardReplaceInsert)
	if err != nil {
		tx.Rollback()
		cutil.Logger.Error(fmt.Sprintf("换卡操作数据库错误%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	tx.Commit()
	return resp, nil
}

// CancelCard 卡注销
func (s *Server) CancelCard(ctx context.Context, request *proto.CancelCardRequest) (*proto.NoDataResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("CancelCard")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}
	metadata := pkgs.GetMetadata(ctx)
	cardID := uuid.FromStringOrNil(request.CardId)
	if cardID == uuid.Nil {
		cutil.Logger.Error(fmt.Sprintf("CancelCard 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	card, err := model.ShowCard(cardID, metadata.MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("注销卡查询卡信息数据库错误%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	if len(request.BankAccount) == 0 || len(request.BankName) == 0 || len(request.MoneyReceiver) == 0 || len(request.Reason) == 0 {
		cutil.Logger.Error(fmt.Sprintf("CancelCard 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	now := time.Now()

	cardCancel := model.TableCardCancel{
		ID:            uuid.NewV4(),
		CardID:        &cardID,
		AccountID:     card.AccountIDs,
		BankAccount:   request.BankAccount,
		BankName:      request.BankName,
		MoneyReceiver: request.MoneyReceiver,
		Reason:        request.Reason,
		ApplyStaffID:  &metadata.StaffID,
		ApplyAt:       &now,
		Status:        model.CardCancelStatusApply,
		MerchantID:    &metadata.MerchantID,
	}
	err = model.CardCancel(cardCancel)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("CardCancel 注销卡数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	return resp, nil
}

// ValidateCardPassword 验证卡密码
func (s *Server) ValidateCardPassword(ctx context.Context, request *proto.ValidateCardPasswordRequest) (*proto.NoDataResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ValidCardPassword")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}

	if request.CardCode == "" || request.Password == "" {
		cutil.Logger.Error(fmt.Sprintf("ValidCardPassword 请求参数错误:%v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	card, err := model.ShowCardByCode(request.CardCode, pkgs.GetMetadata(ctx).MerchantID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("model.ShowCardByCode 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrUnprocessableEntity
		}
		return resp, nil
	}

	if card.Status != model.CardStatusActive {
		cutil.Logger.Error(fmt.Sprintf("ValidCardPassword，卡状态异常：%s", card.Status))
		resp.ErrorMessage = "卡非激活状态"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	if card.Password != "" && card.Password == util.Sha256(request.Password) {
		return resp, nil
	}

	resp.ErrorMessage = "卡密码错误"
	resp.ErrorCode = pkgs.ErrUnprocessableEntity
	return resp, nil

}

// Consume 消费
func (s *Server) Consume(ctx context.Context, req *proto.ConsumeRequest) (*proto.NoDataResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("Consume")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	staffID := pkgs.GetMetadata(ctx).StaffID
	branchID := uuid.FromStringOrNil(req.BranchId)
	cardID := uuid.FromStringOrNil(req.CardId)
	card, err := model.ShowCard(cardID, merchantID)
	pointBillID := uuid.FromStringOrNil(req.PointBillId)
	posBillID := uuid.FromStringOrNil(req.PosBillId)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("Consume 获取卡信息数据库错误:%v", err))
		resp.ErrorMessage = "消费失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	if !req.IgnorePassword && util.Sha256(req.Password) != card.Password {
		resp.ErrorMessage = "卡密码错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	var products, packages fields.ProductPackageTicketArr
	for _, v := range req.Products {
		products = append(products, fields.ProductPackageTicket{
			ID:    uuid.FromStringOrNil(v.Id),
			Count: v.Number,
		})
	}
	for _, v := range req.Packages {
		packages = append(packages, fields.ProductPackageTicket{
			ID:    uuid.FromStringOrNil(v.Id),
			Count: v.Number,
		})
	}

	tx := model.GetDBEntity().Conn.Begin()
	_, errMsg := s.consume(ctx, &consumeRequest{
		Tx:         tx,
		CardID:     cardID,
		AccountIDs: card.AccountIDs.Slice(),
		CostValue:  req.Amount,
		Products:   &products,
		Packages:   &packages,
		BranchID:   branchID,
		StaffID:    staffID,
		MerchantID: merchantID,
		BillType:   string(model.BillTypeConsume),
		PosBillID:  posBillID,
	})
	if errMsg.ErrorCode != pkgs.Success {
		tx.Rollback()
		resp.ErrorMessage = errMsg.ErrorMessage
		resp.ErrorCode = errMsg.ErrorCode
		return resp, nil
	}

	if pointBillID != uuid.Nil {
		extensionResp, err := s.memberExtension().RemoveMemberPointBillLock(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &memberExtension.RemoveMemberPointBillLockRequest{
			BillId:  pointBillID.String(),
			IsValid: true,
		})
		if err != nil || extensionResp == nil || extensionResp.ErrorCode != pkgs.Success {
			tx.Rollback()
			cutil.Logger.Error(fmt.Sprintf("Consume 扣除积分错误:%v, %v", err, extensionResp))
			resp.ErrorMessage = "消费失败"
			resp.ErrorCode = pkgs.ErrInternal
			return resp, nil
		}
	}

	tx.Commit()
	return resp, nil
}

// RefundBill 退款
func (s *Server) RefundBill(ctx context.Context, req *proto.RefundBillRequest) (*proto.NoDataResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("RefundBill")
	resp := &proto.NoDataResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID
	staffID := pkgs.GetMetadata(ctx).StaffID
	id := uuid.FromStringOrNil(req.PosBillId)

	bills, err := model.GetAccountBillsByPosBillID(id, model.BillStatusSuccess)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("GetAccountBillsByPosBillID 获取账单数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "退款失败"
		return resp, nil
	}
	refundBills := make([]model.TableAccountBill, len(bills))
	for i := range bills {
		// 生成账单号
		billCode, err := s.getBillCode(ctx, *bills[i].BranchID, time.Now())
		if err != nil || billCode == "" {
			cutil.Logger.Error(fmt.Sprintf("生成账单号错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "退款失败"
			return resp, nil
		}
		account, err := model.ShowCardAccount(*bills[i].AccountID)
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("获取账户信息错误:%v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "退款失败"
			return resp, nil
		}
		account.BaseValue += bills[i].BaseValue
		account.GiftValue += bills[i].GiftValue
		refundBills[i] = model.TableAccountBill{
			BillCode:       billCode,
			AccountID:      bills[i].AccountID,
			CardID:         bills[i].CardID,
			CardCode:       bills[i].CardCode,
			BranchID:       bills[i].BranchID,
			ChangeValue:    bills[i].ChangeValue,
			ChangeCategory: model.BillCategoryRecharge,
			ChangeType:     model.BillTypeRefund,
			BaseValue:      bills[i].BaseValue,
			GiftValue:      bills[i].GiftValue,
			AfterAccount:   account,
			StaffID:        &staffID,
			MerchantID:     &merchantID,
			BaseValueLeft:  bills[i].BaseValue,
			GiftValueLeft:  bills[i].GiftValue,
			PosBillID:      &id,
			Status:         model.BillStatusSuccess,
		}
	}
	tx := model.GetDBEntity().Conn.Begin()
	if err := model.RefundBill(tx, bills, refundBills); err != nil {
		tx.Rollback()
		cutil.Logger.Error(fmt.Sprintf("退款账单修改数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "退款失败"
		return resp, nil
	}
	extensionResp, err := s.memberExtension().RefundPoint(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &memberExtension.RefundPointRequest{PosBillId: req.PosBillId})
	if err != nil || extensionResp == nil || extensionResp.ErrorCode != pkgs.Success {
		tx.Rollback()
		cutil.Logger.Error(fmt.Sprintf("退积分rpc错误:%v %v", err, extensionResp))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "退款失败"
		return resp, nil
	}
	tx.Commit()

	return resp, nil
}

func toCardProto(card *model.TableCard) *proto.Card {
	openAt := int32(0)
	if card.OpenedAt != nil {
		openAt = int32(card.OpenedAt.Unix())
	}
	return &proto.Card{
		Id:             card.ID.String(),
		Category:       string(card.Category),
		SubCategory:    string(card.SubCategory),
		Code:           card.Code,
		CreateBranchId: cutil.UUIDToString(card.CreateBranchID),
		CreateStaffId:  cutil.UUIDToString(card.CreateStaffID),
		Status:         string(card.Status),
		OpenedAt:       openAt,
		OpenOperator:   cutil.UUIDToString(card.OpenStaffID),
		AccountId:      card.AccountIDs.ToStringArr(),
		OpenBranchId:   cutil.UUIDToString(card.OpenBranchID),
		MemberId:       cutil.UUIDToString(card.MemberID),
		PrimaryId:      cutil.UUIDToString(card.PrimaryID),
	}
}

func validCardStatus(currStatus model.CardStatus, action string) (model.CardStatus, error) {
	err := errors.New("卡状态异常")
	switch action {
	//挂失
	case model.CardActionLost:
		if currStatus == model.CardStatusActive {
			return model.CardStatusLost, nil
		}
		return "", err
	//找回
	case model.CardActionFind:
		if currStatus == model.CardStatusLost {
			return model.CardStatusActive, nil
		}
		return "", err
	//注销申请
	case model.CardActionCancel:
		if currStatus == model.CardStatusActive {
			return model.CardStatusCancelling, nil
		}
		return "", err
	//注销审核
	case model.CardActionCancelExamine:
		if currStatus == model.CardStatusCancelling {
			return model.CardStatusCancelled, nil
		}
		return "", err
	default:
		return "", err
	}
}

func (s *Server) getBillCode(ctx context.Context, branchID uuid.UUID, t time.Time) (string, error) {
	// 查询门店
	areaCode, _ := model.GetAreaCode(branchID)
	if areaCode == "" {
		branchResp, err := s.merchantBasic().ShowBranch(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &merchantBasic.ShowBranchRequest{Id: branchID.String()})
		if err != nil || branchResp == nil || branchResp.ErrorCode != pkgs.Success {
			cutil.Logger.Error(fmt.Sprintf("查询门店错误:%v, %v", err, branchResp))
			return "", errors.New("查询门店错误")
		}

		areaResp, err := s.basic().ShowArea(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &basicProto.ShowAreaRequest{Code: branchResp.Data.CityId})
		if err != nil || areaResp == nil || areaResp.ErrorCode != pkgs.Success {
			cutil.Logger.Error(fmt.Sprintf("查询城市编号错误:%v, %v", err, areaResp))
			return "", errors.New("查询城市编号错误")
		}
		areaCode = areaResp.Data.Telcode
		if err := model.SetAreaCode(branchID, areaCode); err != nil {
			cutil.Logger.Error(fmt.Sprintf("缓存billcode redis错误:%v", err))
		}
	}
	code, err := model.ShowBillCodeSeq()
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("生成账单自增序列号错误:%v", err))
		return "", err
	}
	return fmt.Sprintf("%s%s%06d", areaCode, t.Format("20060102"), code), nil
}

func toProtoBillData(bill model.TableAccountBill) *proto.BillData {
	var cardAccount model.TableCardAccount
	var payments *proto.Payments
	if bill.Payments != nil {
		payments = &proto.Payments{
			Wechat: bill.Payments.Wechat,
			Cash:   bill.Payments.Cash,
			Alipay: bill.Payments.Alipay,
			Card:   bill.Payments.Card,
		}
	}
	if bill.AfterAccount != nil {
		cardAccount = *bill.AfterAccount
	}
	return &proto.BillData{
		BillId:          bill.ID.String(),
		BillCode:        bill.BillCode,
		AccountId:       cutil.UUIDToString(bill.AccountID),
		CardId:          cutil.UUIDToString(bill.CardID),
		CardCode:        bill.CardCode,
		BranchId:        cutil.UUIDToString(bill.BranchID),
		ChangeValue:     bill.ChangeValue,
		ChangeCategory:  string(bill.ChangeCategory),
		ChangeType:      string(bill.ChangeType),
		BaseValue:       bill.BaseValue,
		GiftValue:       bill.GiftValue,
		Payments:        payments,
		StaffId:         cutil.UUIDToString(bill.StaffID),
		OperatorComment: bill.OperatorComment,
		MerchantId:      cutil.UUIDToString(bill.MerchantID),
		CreatedAt:       int32(bill.CreatedAt.Unix()),
		TotalBaseValue:  cardAccount.BaseValue,
		TotalGiftValue:  cardAccount.GiftValue,
	}
}

// 划账前验证账户和卡的状态
func validateTransferAccount(account *model.TableCardAccount) error {
	if account.Status == model.AccountStatusFrozen {
		return errors.New("账户已冻结")
	}
	if account.Status == model.AccountStatusCancelled {
		return errors.New("账户已注销")
	}
	return nil
}

func (s *Server) listPromotionOption(ctx context.Context, promotions []*proto.PromotionCount, rechargeValue int32, branchID uuid.UUID) (value *rechargeMoney, promotionOptions model.PromotionOptions, ppt productPackageTicket, err error) {
	promotionOptions = make(model.PromotionOptions, 0)
	promotionIDs := make([]string, 0)
	promotionMap := make(map[string]int32)
	baseValue := int32(0)
	giftValue := int32(0)
	for _, v := range promotions {
		promotionIDs = append(promotionIDs, v.Id)
		promotionMap[v.Id] = v.Count
	}

	promotionResp, err := s.memberPrivate().ListPromotionOption(pkgs.MetadataContent(pkgs.GetMetadata(ctx)), &private.ListPromotionOptionRequest{Ids: promotionIDs, BranchId: branchID.String()})
	if err != nil || promotionResp == nil {
		cutil.Logger.Error(fmt.Sprintf("查询优惠方案错误:%v, %v", err, promotionResp))
		return nil, nil, ppt, fmt.Errorf("查询优惠方案错误:%v", err)
	}
	var tagID uuid.UUID
	productMap := make(map[uuid.UUID]proto.ProductPackage)
	packageMap := make(map[uuid.UUID]proto.ProductPackage)
	ticketMap := make(map[uuid.UUID]proto.ProductPackage)
	if promotionResp.ErrorCode == pkgs.Success {
		if len(promotionResp.Data.Data) == 0 {
			cutil.Logger.Error("没有优惠方案")
			return nil, nil, ppt, fmt.Errorf("没有优惠方案")
		}
		for _, v := range promotionResp.Data.Data {
			for i := 0; i < int(promotionMap[v.Id]); i++ {
				if rechargeValue < int32(v.RechargeValue) {
					cutil.Logger.Error(fmt.Sprintf("充值金额与优惠方案不匹配,充值金额:%v,优惠方案金额:%v", rechargeValue, v.RechargeValue))
					return nil, nil, ppt, fmt.Errorf("充值金额与优惠方案不匹配")
				}
				rechargeValue -= int32(v.RechargeValue)
				baseValue += int32(v.BaseValue)
				giftValue += int32(v.GiftValue)
				if tagID == uuid.Nil {
					if tagID = uuid.FromStringOrNil(v.TagId); tagID == uuid.Nil {
						cutil.Logger.Error("优惠方案标签ID为nil")
						return nil, nil, ppt, fmt.Errorf("优惠方案标签ID为nil")
					}
				} else if tagID != uuid.FromStringOrNil(v.TagId) {
					cutil.Logger.Error("优惠方案标签ID不同")
					return nil, nil, ppt, fmt.Errorf("优惠方案标签ID不同")
				}
			}
			getProductPackageTicketMap(productMap, v.Products, promotionMap[v.Id])
			getProductPackageTicketMap(packageMap, v.Packages, promotionMap[v.Id])
			getProductPackageTicketMap(ticketMap, v.Tickets, promotionMap[v.Id])

			// 账单记录优惠详情
			var products, packages, tickets *fields.ProductPackageTicketArr
			if len(v.Products) != 0 {
				products = new(fields.ProductPackageTicketArr)
				for j := range v.Products {
					*products = append(*products, fields.ProductPackageTicket{
						ID:    uuid.FromStringOrNil(v.Products[j].Id),
						Code:  v.Products[j].Code,
						Count: v.Products[j].Number,
						Price: v.Products[j].Price,
						Title: v.Products[j].Title,
					})
				}
			}
			if len(v.Packages) != 0 {
				packages = new(fields.ProductPackageTicketArr)
				for j := range v.Packages {
					*packages = append(*packages, fields.ProductPackageTicket{
						ID:    uuid.FromStringOrNil(v.Packages[j].Id),
						Code:  v.Packages[j].Code,
						Count: v.Packages[j].Number,
						Price: v.Packages[j].Price,
						Title: v.Packages[j].Title,
					})
				}
			}
			if len(v.Tickets) != 0 {
				tickets = new(fields.ProductPackageTicketArr)
				for j := range v.Tickets {
					*tickets = append(*tickets, fields.ProductPackageTicket{
						ID:    uuid.FromStringOrNil(v.Tickets[j].Id),
						Code:  v.Tickets[j].Code,
						Count: v.Tickets[j].Number,
						Price: v.Tickets[j].Price,
						Title: v.Tickets[j].Title,
					})
				}
			}
			promotionOptions = append(promotionOptions, model.PromotionOption{
				ID:            uuid.FromStringOrNil(v.Id),
				Name:          v.Name,
				Count:         int(promotionMap[v.Id]),
				RechargeValue: int32(v.RechargeValue),
				BaseValue:     int32(v.BaseValue),
				GiftValue:     int32(v.GiftValue),
				Products:      products,
				Packages:      packages,
				Tickets:       tickets,
			})
		}
	} else {
		cutil.Logger.Error(fmt.Sprintf("查询优惠方案错误:%v", promotionResp))
		return nil, nil, ppt, fmt.Errorf("查询优惠方案错误:%v", promotionResp)
	}
	ppt.products, ppt.packages, ppt.tickets = getProductPackageTicketArr(productMap), getProductPackageTicketArr(packageMap), getProductPackageTicketArr(ticketMap)
	baseValue += rechargeValue
	value = &rechargeMoney{baseValue, giftValue, tagID}
	return value, promotionOptions, ppt, nil
}

func (s *Server) SearchCards(ctx context.Context, request *proto.SearchCardsRequest) (*proto.SearchCardsResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SearchCards")
	resp := &proto.SearchCardsResponse{
		ErrorCode: pkgs.Success,
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID

	scopes := []func(*gorm.DB) *gorm.DB{
		cutil.ColumnEqualScope("merchant_id", merchantID),
		cutil.ColumnEqualScope("status", model.CardStatusActive),
		cutil.ColumnLikeScope("code", request.CardCode),
	}

	cards, err := model.SearchCards(scopes)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("SearchCards 数据库错误:%+v", err))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var protoData []*proto.CardWithAccount
	if len(cards) > 0 {
		protoData, err = buildMemberAccounts(&cards)
		if err != nil {
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = "账户查询失败"
			return resp, nil
		}
	}
	resp.Data = protoData
	return resp, nil
}

func getProductPackageTicketMap(m map[uuid.UUID]proto.ProductPackage, pp []*private.ProductPackage, count int32) {
	for _, product := range pp {
		// 总次数 = 一次优惠方案送的 商品/套餐/优惠券 次数*方案的次数 + map存的同一id的次数
		num := product.Number*count + m[uuid.FromStringOrNil(product.Id)].Number
		m[uuid.FromStringOrNil(product.Id)] = proto.ProductPackage{
			Id:     product.Id,
			Code:   product.Code,
			Number: num,
			Price:  product.Price,
			Title:  product.Title,
			Left:   num,
		}
	}
}

func getProductPackageTicketArr(m map[uuid.UUID]proto.ProductPackage) []*proto.ProductPackage {
	pp := make([]*proto.ProductPackage, 0)
	for k := range m {
		v := m[k]
		pp = append(pp, &v)
	}
	return pp
}

func rechargeCardValidate(cardID, merchantID uuid.UUID) (*model.TableCard, *proto.RechargeCardResponse) {
	resp := &proto.RechargeCardResponse{ErrorCode: pkgs.Success}
	card, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			return nil, resp
		}
		cutil.Logger.Error(fmt.Sprintf("RechargeCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "充值失败"
		resp.ErrorCode = pkgs.ErrInternal
		return nil, resp
	}
	if card.Category != model.CardCategoryMember || card.SubCategory != model.CardSubCategoryPrimary {
		cutil.Logger.Error(fmt.Sprintf("RechargeCard 非主卡不能充值:%s", card.Code))
		resp.ErrorMessage = "非主卡不能充值"
		resp.ErrorCode = pkgs.ErrNotFound
		return nil, resp
	}
	if card.Status != model.CardStatusActive {
		cutil.Logger.Error(fmt.Sprintf("RechargeCard 卡状态错误:%s", card.Status))
		resp.ErrorMessage = "卡状态错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return nil, resp
	}
	return card, resp
}

func activeSecondaryCardValidate(cardID, primaryID, primaryAccountID, merchantID uuid.UUID, password string, verified bool) (*model.TableCard, *model.TableCard, *proto.ActiveSecondaryCardResponse) {
	resp := &proto.ActiveSecondaryCardResponse{ErrorCode: pkgs.Success}
	tabCard, err := model.ShowCard(cardID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 卡不存在:%v", cardID))
			return nil, nil, resp
		}
		cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "激活卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return nil, nil, resp
	}

	if tabCard.Status != model.CardStatusInit {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "卡已激活"
		return nil, nil, resp
	}

	primaryCard, err := model.ShowCard(primaryID, merchantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorMessage = "卡不存在"
			resp.ErrorCode = pkgs.ErrNotFound
			cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 主卡不存在:%v", primaryID.String()))
			return nil, nil, resp
		}
		cutil.Logger.Error(fmt.Sprintf("ActiveSecondaryCard 查询卡数据库错误:%v", err))
		resp.ErrorMessage = "激活卡失败"
		resp.ErrorCode = pkgs.ErrInternal
		return nil, nil, resp
	}
	if primaryCard.Status != model.CardStatusActive {
		resp.ErrorMessage = "主卡状态错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return nil, nil, resp
	}
	//为副卡时，验证主卡密码，设置主卡id
	if verified {
		if primaryCard.Password != password {
			resp.ErrorMessage = "主卡密码错误"
			resp.ErrorCode = pkgs.ErrPassword
			return nil, nil, resp
		}
	}
	if !cutil.ArrContainElement(primaryCard.AccountIDs.Slice(), primaryAccountID) {
		resp.ErrorMessage = "主卡与账户不匹配"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return nil, nil, resp
	}
	return tabCard, primaryCard, resp
}

func activePrimaryCardData(memberID, branchID, merchantID, operatorID uuid.UUID, card *model.TableCard, password string, money *rechargeMoney) (*model.TableCardAccount, error) {
	accounts, err := model.GetAccountByMemberID(memberID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ActivePrimaryCard 获取会员账户错误:%v", err))
		return nil, err
	}

	var cardAccount model.TableCardAccount
	for _, v := range accounts {
		if *v.TagID == money.tagID && v.Category == model.CardSubCategoryPrimary {
			if v.Status != model.AccountStatusActivated {
				return nil, model.ErrAccountAbnormal
			}
			cardAccount = v
			break
		}
	}

	if cardAccount.ID == uuid.Nil {
		cardAccount = model.TableCardAccount{
			ID:         uuid.NewV4(),
			MemberID:   &memberID,
			BranchID:   &branchID,
			Status:     model.AccountStatusActivated,
			MerchantID: merchantID,
			TagID:      &money.tagID,
			Category:   model.CardSubCategoryPrimary,
		}
	}
	cardAccount.BaseValue += money.baseValue
	cardAccount.GiftValue += money.giftValue
	now := time.Now()
	card.Category = model.CardCategoryMember
	card.SubCategory = model.CardSubCategoryPrimary
	card.OpenBranchID = &branchID
	if !cutil.ArrContainElement(card.AccountIDs.Slice(), cardAccount.ID) {
		if card.AccountIDs == nil {
			card.AccountIDs = &fields.UUIDArr{}
		}
		*card.AccountIDs = append(*card.AccountIDs, cardAccount.ID)
	}
	card.Status = model.CardStatusActive
	card.OpenedAt = &now
	card.OpenStaffID = &operatorID
	if password == "" {
		card.Password = util.DefaultCardPassword
	} else {
		card.Password = password
	}
	card.MemberID = &memberID
	return &cardAccount, nil
}

func accountProductPackage(account *model.TableCardAccount, ppt productPackageTicket) {
	if len(ppt.products) != 0 {
		if account.Products != nil {
			for i := range ppt.products {
				have := false
				for j := range *account.Products {
					if (*account.Products)[i].ID.String() == ppt.products[j].Id {
						(*account.Products)[i].Count += ppt.products[j].Number
						have = true
						break
					}
				}
				if !have {
					*account.Products = append(*account.Products, fields.ProductPackageTicket{
						ID:    uuid.FromStringOrNil(ppt.products[i].Id),
						Code:  ppt.products[i].Code,
						Count: ppt.products[i].Number,
						Title: ppt.products[i].Title,
					})
				}
			}
		} else {
			for i := range ppt.products {
				*account.Products = append(*account.Products, fields.ProductPackageTicket{
					ID:    uuid.FromStringOrNil(ppt.products[i].Id),
					Code:  ppt.products[i].Code,
					Count: ppt.products[i].Number,
					Title: ppt.products[i].Title,
				})
			}
		}
	}

	if len(ppt.packages) != 0 {
		if account.Packages != nil {
			for i := range ppt.packages {
				have := false
				for j := range *account.Packages {
					if (*account.Packages)[i].ID.String() == ppt.packages[j].Id {
						(*account.Packages)[i].Count += ppt.packages[j].Number
						have = true
						break
					}
				}
				if !have {
					*account.Packages = append(*account.Packages, fields.ProductPackageTicket{
						ID:    uuid.FromStringOrNil(ppt.packages[i].Id),
						Code:  ppt.packages[i].Code,
						Count: ppt.packages[i].Number,
						Title: ppt.packages[i].Title,
					})
				}
			}
		} else {
			for i := range ppt.packages {
				*account.Packages = append(*account.Packages, fields.ProductPackageTicket{
					ID:    uuid.FromStringOrNil(ppt.packages[i].Id),
					Code:  ppt.packages[i].Code,
					Count: ppt.packages[i].Number,
					Title: ppt.packages[i].Title,
				})
			}
		}
	}
}

type rechargeMoney struct {
	baseValue int32
	giftValue int32
	tagID     uuid.UUID
}

type productPackageTicket struct {
	products []*proto.ProductPackage
	packages []*proto.ProductPackage
	tickets  []*proto.ProductPackage
}
