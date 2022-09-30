package fields

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// UUIDArr 数组类型的UUID 集合
type UUIDArr []uuid.UUID

// StringArrToUUIDArr 切片数组转化
func StringArrToUUIDArr(val []string) (UUIDArr, error) {
	res := make(UUIDArr, 0)
	for _, v := range val {
		if len(v) == 0 {
			continue
		}

		i, err := uuid.FromString(v)
		if err != nil {
			return res, fmt.Errorf("转化uuid错误,:%s", v)
		}
		res = append(res, i)
	}
	return res, nil
}

// Value 转换为数据库字段类型
func (ua UUIDArr) Value() (driver.Value, error) {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	for k, u := range ua {
		if k > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(u.String())
	}
	buffer.WriteString("}")

	return buffer.String(), nil
}

// Slice slice
func (ua *UUIDArr) Slice() []uuid.UUID {
	if ua != nil {

		var response []uuid.UUID

		for _, id := range *ua {
			response = append(response, id)
		}

		return response
	}
	return nil
}

// Scan 将数据库数据映射到结构体
func (ua *UUIDArr) Scan(src interface{}) error {
	if nil == src {
		return nil
	}

	val := src.(string)
	value := strings.Trim(strings.Trim(val, "{"), "}")
	if len(value) == 0 {
		return nil
	}

	var res UUIDArr
	for _, v := range strings.Split(value, ",") {
		if len(v) == 0 {
			continue
		}
		res = append(res, uuid.FromStringOrNil(v))
	}

	*ua = res

	return nil
}

// ToStringArr 转化为[]string
func (ua *UUIDArr) ToStringArr() []string {
	var res []string
	if ua != nil {
		for _, v := range *ua {
			res = append(res, v.String())
		}
	}
	return res
}

// ToMetadataString 转化为metadata string
func (ua UUIDArr) ToMetadataString() string {
	var buffer bytes.Buffer

	for k, u := range ua {
		if k > 0 {
			buffer.WriteString("|")
		}
		buffer.WriteString(u.String())
	}

	return buffer.String()
}

// MetadataStringToUUIDArr metadata to uuid arr
func MetadataStringToUUIDArr(val string) (UUIDArr, error) {
	var res UUIDArr
	for _, v := range strings.Split(val, "|") {
		if len(v) == 0 {
			continue
		}

		i, err := uuid.FromString(v)
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}
	return res, nil
}
