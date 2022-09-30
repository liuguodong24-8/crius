package pkgs

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

// Params map[string]interface{}
type Params map[string]interface{}

// ParamsArr []map[string]interface{}
type ParamsArr []map[string]interface{}

// Set 设置值
func (p Params) Set(k string, v interface{}) Params {
	p[k] = v

	return p
}

// Get 获取值
func (p Params) Get(k string) (v interface{}) {
	if v, ok := p[k]; ok {
		return v
	}

	return nil
}

// Exists 判断是否存在
func (p Params) Exists(k string) bool {
	_, ok := p[k]

	return ok
}

// JSON 获取json
func (p Params) JSON() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(j)
}

// Value 返回数据库可识别类型
func (p Params) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *Params) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("util params value")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("util params value error")
	}

	return nil
}

// MakeParams 生成params
func MakeParams(data interface{}) Params {
	p := make(Params)

	j, e := json.Marshal(data)
	if e != nil {
		return p
	}

	if err := json.Unmarshal(j, &p); nil != err {
		return Params{}
	}

	return p
}

// Value 返回数据库可识别类型
func (p ParamsArr) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *ParamsArr) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("util params value")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("util params value error")
	}

	return nil
}

// JSON 获取json
func (p *ParamsArr) JSON() string {
	if p != nil {
		j, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}

		return string(j)
	}
	return ""
}

// MakeParamsArr 生成paramsArr
func MakeParamsArr(data interface{}) ParamsArr {
	p := make(ParamsArr, 0)

	j, e := json.Marshal(data)
	if e != nil {
		return p
	}

	if err := json.Unmarshal(j, &p); nil != err {
		return ParamsArr{}
	}

	return p
}
