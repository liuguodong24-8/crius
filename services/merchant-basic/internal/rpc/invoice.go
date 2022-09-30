package rpc

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
)

// CreateInvoice 保存票据
func (s *Server) CreateInvoice(ctx context.Context, request *proto.CreateInvoiceRequest) (*proto.CreateInvoiceResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("CreateInvoice params", logger.MakeFields(request)).Info("CreateInvoice")
	resp := &proto.CreateInvoiceResponse{
		ErrorCode: pkgs.Success,
	}

	if request.Action == "" || request.InvoiceData == "" {
		crius.Logger.Error(fmt.Sprintf("CreateInvoice 请求参数错误, %v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "请求参数错误"
		return resp, nil
	}

	merchantID := pkgs.GetMetadata(ctx).MerchantID

	if err := validateInvoiceAction(request.Action); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateInvoice 请求参数错误, %v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	invoice := &model.TableInvoice{
		ID:         uuid.NewV4(),
		Action:     request.Action,
		Data:       request.InvoiceData,
		MerchantID: &merchantID,
	}
	if err := model.CreateInvoice(invoice); err != nil {
		crius.Logger.Error(fmt.Sprintf("CreateInvoice 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "数据库错误"
		return resp, nil
	}
	return resp, nil
}

func validateInvoiceAction(action string) error {
	actionMap := map[string]bool{
		model.InvoiceActionOpenCard: true,
		model.InvoiceActionRecharge: true,
	}

	if !actionMap[action] {
		return errors.New("非法的票据场景值")
	}

	return nil
}
