package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
	"gorm.io/gorm"
)

// GetBills 账单列表
func (s *Server) GetBills(ctx context.Context, request *proto.GetBillsRequest) (*proto.GetBillsResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("GetBills")
	resp := &proto.GetBillsResponse{
		ErrorCode: pkgs.Success,
	}

	accountID := uuid.FromStringOrNil(request.AccountId)
	if accountID == uuid.Nil {
		cutil.Logger.Error(fmt.Sprintf("GetBills 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	count, err := model.CountBills(accountID, model.BillCategory(request.Category))
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("CountBills 账单查询数据库错误:%v", err))
		resp.ErrorMessage = "查询失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	if count == 0 {
		resp.Data = &proto.BillsData{Total: 0}
		return resp, nil
	}

	bills, err := model.GetBills(accountID, model.BillCategory(request.Category), request.Offset, request.Limit)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("GetBills 账单查询数据库错误:%v", err))
		resp.ErrorMessage = "查询失败"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}
	var billData []*proto.Bill
	if len(bills) > 0 {
		for _, bill := range bills {
			pp, err := model.GetProductPackagesByBillID(bill.ID)
			if err != nil {
				cutil.Logger.Error(fmt.Sprintf("GetBills 账单查询获取product_package数据库错误:%v", err))
				resp.ErrorMessage = "查询失败"
				resp.ErrorCode = pkgs.ErrInternal
				return resp, nil
			}
			protoBill := toBillProto(&bill)
			protoBill.Products, protoBill.Packages = toProtoBillProductPackage(pp)
			billData = append(billData, protoBill)
		}
	}
	resp.Data = &proto.BillsData{Bills: billData, Total: int32(count)}
	return resp, nil
}

// ShowBill 通过ID查询账单信息
func (s *Server) ShowBill(ctx context.Context, request *proto.ShowBillRequest) (*proto.ShowBillResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowBill")
	resp := &proto.ShowBillResponse{
		ErrorCode: pkgs.Success,
	}

	billID := uuid.FromStringOrNil(request.BillId)
	if billID == uuid.Nil {
		cutil.Logger.Error(fmt.Sprintf("ShowBill 请求参数错误:%+v", request))
		resp.ErrorMessage = "请求参数错误"
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		return resp, nil
	}

	bill, err := model.ShowBill(billID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ShowBill 数据库错误:%+v", request))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorMessage = "数据不存在"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}
	pp, err := model.GetProductPackagesByBillID(billID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ShowBill 查询product_package数据库错误:%+v", request))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = toBillProto(bill)
	resp.Data.Products, resp.Data.Packages = toProtoBillProductPackage(pp)
	return resp, nil
}

//ShowBillByCode 通过账单号查询信息
func (s *Server) ShowBillByCode(ctx context.Context, request *proto.ShowBillByCodeRequest) (*proto.ShowBillByCodeResponse, error) {
	defer cutil.CatchException()
	cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowBillByCode")
	resp := &proto.ShowBillByCodeResponse{
		ErrorCode: pkgs.Success,
	}

	bill, err := model.ShowBillByCode(request.BillCode)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ShowBillByCode 数据库错误:%+v", request))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.ErrorMessage = "数据不存在"
			resp.ErrorCode = pkgs.ErrNotFound
		}
		return resp, nil
	}
	pp, err := model.GetProductPackagesByBillID(bill.ID)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("ShowBillByCode 查询product_package数据库错误:%+v", request))
		resp.ErrorMessage = "数据库错误"
		resp.ErrorCode = pkgs.ErrInternal
		return resp, nil
	}

	resp.Data = toBillProto(bill)
	resp.Data.Products, resp.Data.Packages = toProtoBillProductPackage(pp)
	return resp, nil
}

func toBillProto(bill *model.TableAccountBill) *proto.Bill {
	var account model.TableCardAccount
	if bill.AfterAccount != nil {
		account = *bill.AfterAccount
	}
	var paymentsProto *proto.Payments
	if bill.Payments != nil {
		paymentsProto = &proto.Payments{
			Wechat: bill.Payments.Wechat,
			Cash:   bill.Payments.Cash,
			Alipay: bill.Payments.Alipay,
			Card:   bill.Payments.Card,
		}
	}

	return &proto.Bill{
		Id:             bill.ID.String(),
		BillCode:       bill.BillCode,
		CreateAt:       int32(bill.CreatedAt.Unix()),
		Branch:         cutil.UUIDToString(bill.BranchID),
		CostBaseValue:  bill.BaseValue,
		CostGiftValue:  bill.GiftValue,
		CurrBaseValue:  account.BaseValue,
		CurrGiftValue:  account.GiftValue,
		ChangeCategory: toBillCategoryProto(bill.ChangeCategory, bill.ChangeType),
		Payments:       paymentsProto,
		StaffId:        cutil.UUIDToString(bill.StaffID),
		BaseValueLeft:  bill.BaseValueLeft,
		GiftValueLeft:  bill.GiftValueLeft,
		CardId:         cutil.UUIDToString(bill.CardID),
		CardCode:       bill.CardCode,
	}
}

func toProtoBillProductPackage(pp []model.TableProductPackage) (products []*proto.ProductPackage, packages []*proto.ProductPackage) {
	for i := range pp {
		if pp[i].Category == model.ProductPackageCategoryProduct {
			products = append(products, &proto.ProductPackage{
				Id:               pp[i].ID.String(),
				ProductPackageId: pp[i].ProductPackageID.String(),
				Code:             pp[i].Code,
				Number:           pp[i].Number,
				Price:            pp[i].Price,
				Title:            pp[i].Title,
				Left:             pp[i].Left,
				Category:         string(pp[i].Category),
			})
		}
		if pp[i].Category == model.ProductPackageCategoryPackage {
			packages = append(packages, &proto.ProductPackage{
				Id:               pp[i].ID.String(),
				ProductPackageId: pp[i].ProductPackageID.String(),
				Code:             pp[i].Code,
				Number:           pp[i].Number,
				Price:            pp[i].Price,
				Title:            pp[i].Title,
				Left:             pp[i].Left,
				Category:         string(pp[i].Category),
			})
		}
	}
	return
}

func toBillCategoryProto(billCategory model.BillCategory, billType model.BillType) string {
	if billCategory == model.BillCategoryConsume && billType == model.BillTypeDeduction {
		return string(model.BillTypeDeduction)
	}
	return string(billCategory)
}
