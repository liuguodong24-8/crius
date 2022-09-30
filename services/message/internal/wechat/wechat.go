package wechat

// CommonErr 微信统一返回错误码
type CommonErr struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
