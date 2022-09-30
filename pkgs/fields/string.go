package fields

import (
	"bytes"
	"database/sql/driver"
	"strings"
)

// StringArr 数组类型的String 集合
type StringArr []string

// Value 转换为数据库字段类型
func (sa StringArr) Value() (driver.Value, error) {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	for k, s := range sa {
		if k > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(s)
	}
	buffer.WriteString("}")

	return buffer.String(), nil
}

// Slice slice
func (sa *StringArr) Slice() []string {
	if sa != nil {
		var response []string
		for _, s := range *sa {
			response = append(response, s)
		}

		return response
	}
	return nil
}

// Scan 将数据库数据映射到结构体
func (sa *StringArr) Scan(src interface{}) error {
	if nil == src {
		return nil
	}

	val := src.(string)
	value := strings.Trim(strings.Trim(val, "{"), "}")
	if len(value) == 0 {
		return nil
	}

	var res StringArr
	for _, v := range strings.Split(value, ",") {
		if len(v) == 0 {
			continue
		}
		res = append(res, v)
	}

	*sa = res

	return nil
}

// ToInterfaceArr 转换为interface数组
func (sa StringArr) ToInterfaceArr() (result []interface{}) {
	for _, v := range sa {
		result = append(result, interface{}(v))
	}
	return
}
