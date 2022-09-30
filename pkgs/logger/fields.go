package logger

import (
	"encoding/json"

	util "gitlab.omytech.com.cn/micro-service/Crius/pkgs"
)

// Fields 日志使用
type Fields util.Params

// MakeFields 生成Fields
func MakeFields(data interface{}) Fields {
	p := make(Fields)

	j, e := json.Marshal(data)
	if e != nil {
		return p
	}

	if err := json.Unmarshal(j, &p); nil != err {
		return Fields{}
	}

	return p
}
