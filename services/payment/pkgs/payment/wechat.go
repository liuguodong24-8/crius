package payment

import (
	"context"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/payment/proto"
)

// WechatUnifiedOrderRequest 下单参数
type WechatUnifiedOrderRequest struct {
	BranchID  string
	StringMap map[string]string
	Int64Map  map[string]int64
}

// WechatUnifiedOrderResponseData 下单返回信息
type WechatUnifiedOrderResponseData struct {
	WechatTradeID  string
	WechatResponse *WechatResponse
}

// WechatUnifiedOrderResponse 下单返回
type WechatUnifiedOrderResponse struct {
	ErrorCode    int32
	ErrorMessage string
	Data         *WechatUnifiedOrderResponseData
}

// WechatResponse 微信返回信息
type WechatResponse struct {
	ReturnCode string
	ReturnMsg  string
	Detail     string
	Data       map[string]string
}

// WechatOrderQueryRequest 查询订单状态
type WechatOrderQueryRequest struct {
	BranchID      string
	WechatTradeID string
}

// WechatOrderQueryResponse 查询订单返回
type WechatOrderQueryResponse struct {
	ErrorCode    int32
	ErrorMessage string
	Data         *WechatResponse
}

// WechatCloseOrderRequest 关闭订单
type WechatCloseOrderRequest struct {
	BranchID      string
	WechatTradeID string
}

// WechatCloseOrderResponse 关闭订单返回
type WechatCloseOrderResponse struct {
	ErrorCode    int32
	ErrorMessage string
	Data         *WechatResponse
}

// WechatRefundResponseData 退款返回
type WechatRefundResponseData struct {
	WechatRefundID string
	WechatResponse *WechatResponse
}

// WechatRefundRequest 退款
type WechatRefundRequest struct {
	BranchID      string
	WechatTradeID string
	StringMap     map[string]string
	Int64Map      map[string]int64
}

// WechatRefundResponse 退款返回
type WechatRefundResponse struct {
	ErrorCode    int32
	ErrorMessage string
	Data         *WechatRefundResponseData
}

// WechatRefundQueryRequest 查询退款状态
type WechatRefundQueryRequest struct {
	BranchID       string
	WechatRefundID string
}

// WechatRefundQueryResponse 查询退款返回
type WechatRefundQueryResponse struct {
	ErrorCode    int32
	ErrorMessage string
	Data         *WechatResponse
}

// WechatUnifiedOrder 统一下单
func (c *Client) WechatUnifiedOrder(ctx context.Context, req WechatUnifiedOrderRequest) (WechatUnifiedOrderResponse, error) {
	res, err := c.client.WechatUnifiedOrder(ctx, &proto.WechatUnifiedOrderRequest{
		BranchId:  req.BranchID,
		StringMap: req.StringMap,
		Int64Map:  req.Int64Map,
	})
	if err != nil {
		return WechatUnifiedOrderResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}

	data := &WechatUnifiedOrderResponseData{}

	if res.Data != nil {
		data.WechatTradeID = res.Data.WechatTradeId

		wechatResponse := WechatResponse{}
		if res.Data.WechatResponse != nil {
			wechatResponse.ReturnCode = res.Data.WechatResponse.ReturnCode
			wechatResponse.ReturnMsg = res.Data.WechatResponse.ReturnMsg
			wechatResponse.Detail = res.Data.WechatResponse.Detail
			wechatResponse.Data = res.Data.WechatResponse.Data
		}
	}

	return WechatUnifiedOrderResponse{
		ErrorCode:    res.ErrorCode,
		ErrorMessage: res.ErrorMessage,
		Data:         data,
	}, nil
}

// WechatOrderQuery 订单查询
func (c *Client) WechatOrderQuery(ctx context.Context, req WechatOrderQueryRequest) (WechatOrderQueryResponse, error) {
	res, err := c.client.WechatOrderQuery(ctx, &proto.WechatOrderQueryRequest{
		BranchId:      req.BranchID,
		WechatTradeId: req.WechatTradeID,
	})
	if err != nil {
		return WechatOrderQueryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}

	data := &WechatResponse{}

	if res.Data != nil {
		data.ReturnCode = res.Data.ReturnCode
		data.ReturnMsg = res.Data.ReturnMsg
		data.Detail = res.Data.Detail
		data.Data = res.Data.Data
	}

	return WechatOrderQueryResponse{
		ErrorCode:    res.ErrorCode,
		ErrorMessage: res.ErrorMessage,
		Data:         data,
	}, nil
}

// WechatCloseOrder 订单关闭
func (c *Client) WechatCloseOrder(ctx context.Context, req WechatCloseOrderRequest) (WechatCloseOrderResponse, error) {
	res, err := c.client.WechatCloseOrder(ctx, &proto.WechatCloseOrderRequest{
		BranchId:      req.BranchID,
		WechatTradeId: req.WechatTradeID,
	})
	if err != nil {
		return WechatCloseOrderResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}

	data := &WechatResponse{}

	if res.Data != nil {
		data.ReturnCode = res.Data.ReturnCode
		data.ReturnMsg = res.Data.ReturnMsg
		data.Detail = res.Data.Detail
		data.Data = res.Data.Data
	}

	return WechatCloseOrderResponse{
		ErrorCode:    res.ErrorCode,
		ErrorMessage: res.ErrorMessage,
		Data:         data,
	}, nil
}

// WechatRefund 订单退款
func (c *Client) WechatRefund(ctx context.Context, req WechatRefundRequest) (WechatRefundResponse, error) {
	res, err := c.client.WechatRefund(ctx, &proto.WechatRefundRequest{
		BranchId:      req.BranchID,
		WechatTradeId: req.WechatTradeID,
		StringMap:     req.StringMap,
		Int64Map:      req.Int64Map,
	})
	if err != nil {
		return WechatRefundResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}

	data := &WechatRefundResponseData{}
	if res.Data != nil {
		data.WechatRefundID = res.Data.WechatRefundId
		wechatResponse := WechatResponse{}
		if res.Data.WechatResponse != nil {
			wechatResponse.ReturnCode = res.Data.WechatResponse.ReturnCode
			wechatResponse.ReturnMsg = res.Data.WechatResponse.ReturnMsg
			wechatResponse.Detail = res.Data.WechatResponse.Detail
			wechatResponse.Data = res.Data.WechatResponse.Data
		}
	}

	return WechatRefundResponse{
		ErrorCode:    res.ErrorCode,
		ErrorMessage: res.ErrorMessage,
		Data:         data,
	}, nil
}

// WechatRefundQuery 退款查询
func (c *Client) WechatRefundQuery(ctx context.Context, req WechatRefundQueryRequest) (WechatRefundQueryResponse, error) {
	res, err := c.client.WechatRefundQuery(ctx, &proto.WechatRefundQueryRequest{
		BranchId:       req.BranchID,
		WechatRefundId: req.WechatRefundID,
	})
	if err != nil {
		return WechatRefundQueryResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: err.Error(),
		}, nil
	}

	data := &WechatResponse{}

	if res.Data != nil {
		data.ReturnCode = res.Data.ReturnCode
		data.ReturnMsg = res.Data.ReturnMsg
		data.Detail = res.Data.Detail
		data.Data = res.Data.Data
	}

	return WechatRefundQueryResponse{
		ErrorCode:    res.ErrorCode,
		ErrorMessage: res.ErrorMessage,
		Data:         data,
	}, nil
}
