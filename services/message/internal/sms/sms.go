package sms

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// 强制要求该接口方法必须实现
var _ Interface = (*Entity)(nil)

// SendRequest 请求发送短信参数
type SendRequest struct {
	Sign     string
	AreaCode string
	Mobile   string
	Message  string
}

// SendResponse 返回结构体
type SendResponse struct {
	Code   int         `json:"code"` //0 success
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
	Detail interface{} `json:"detail"`
}

// Interface log interface
type Interface interface {
	Send(ctx context.Context, req SendRequest) (*SendResponse, error)
}

// Config 短信配置
type Config struct {
	URL          string
	SecurityCode string
}

// Entity 日志实例
type Entity struct {
	config Config
}

// URL chunk短信url
const URL = "http://ht.chunk.com.cn:5861/webservice.asmx/SendCustom"

// SecurityCode chunk短信签章
const SecurityCode = "d6f56f1f51a41deba6ffc5"

// NewSms 实例化 sms
func NewSms() (*Entity, error) {
	return &Entity{config: Config{
		URL:          URL,
		SecurityCode: SecurityCode,
	}}, nil
}

// Send 发送短信
func (e *Entity) Send(ctx context.Context, req SendRequest) (*SendResponse, error) {
	values := e.makeSmsValues(req)
	response, err := http.PostForm(e.config.URL, values)
	if err != nil {
		return nil, fmt.Errorf("请求错误:%s", err.Error())
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求错误状态码错误：%d", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("解析返回结果错误:%s", err.Error())
	}

	result, err := parseSmsXML(body)
	if err != nil {
		return nil, fmt.Errorf("解析返回XML错误:%s", err.Error())
	}

	return result, nil
}

func (e *Entity) makeSmsValues(req SendRequest) url.Values {
	attributes := url.Values{}
	attributes.Set("SecurityCode", e.config.SecurityCode)
	attributes.Set("Depart", "")
	attributes.Set("Program", "")
	attributes.Set("Mobile", fmt.Sprintf("%s%s", req.AreaCode, req.Mobile))
	attributes.Set("Message", req.Message)
	attributes.Set("Signature", req.Sign)

	return attributes
}

// XMLContent xml解析格式
type XMLContent struct {
	XMLName xml.Name `xml:"string"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
}

// parseSmsXML 解析xml结构体
func parseSmsXML(bytes []byte) (*SendResponse, error) {
	var content XMLContent

	err := xml.Unmarshal(bytes, &content)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var message SendResponse

	if err := json.Unmarshal([]byte(content.Text), &message); err != nil {
		return nil, errors.WithStack(err)
	}

	return &message, nil
}
