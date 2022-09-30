package fields

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"strings"
)

// Int8Arr 数组类型的int8 集合
type Int8Arr []int8

// Value 转换为数据库字段类型
func (ia Int8Arr) Value() (driver.Value, error) {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	for k, i := range ia {
		if k > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(strconv.Itoa(int(i)))
	}
	buffer.WriteString("}")

	return buffer.String(), nil
}

// Slice slice
func (ia *Int8Arr) Slice() []int8 {
	if ia != nil {
		var response []int8
		for _, i := range *ia {
			response = append(response, i)
		}

		return response
	}
	return nil
}

// SliceInt32 slice
func (ia *Int8Arr) SliceInt32() []int32 {
	if ia != nil {
		var response []int32
		for _, i := range *ia {
			response = append(response, int32(i))
		}

		return response
	}
	return nil
}

// Scan 将数据库数据映射到结构体
func (ia *Int8Arr) Scan(src interface{}) error {
	if nil == src {
		return nil
	}

	val := src.(string)
	value := strings.Trim(strings.Trim(val, "{"), "}")
	if len(value) == 0 {
		return nil
	}

	var res Int8Arr
	for _, v := range strings.Split(value, ",") {
		if len(v) == 0 {
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		res = append(res, int8(i))
	}

	*ia = res

	return nil
}

// IntArrToInt8Arr 转换int数组为Int8Arr
func IntArrToInt8Arr(arr interface{}) Int8Arr {
	var int8Arr Int8Arr
	switch arr.(type) {
	case []int8:
		int8Arr = arr.([]int8)
	case []int16:
		for _, v := range arr.([]int16) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []int32:
		for _, v := range arr.([]int32) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []int64:
		for _, v := range arr.([]int64) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []int:
		for _, v := range arr.([]int) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []uint8:
		for _, v := range arr.([]uint8) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []uint16:
		for _, v := range arr.([]uint16) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []uint32:
		for _, v := range arr.([]uint32) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []uint64:
		for _, v := range arr.([]uint64) {
			int8Arr = append(int8Arr, int8(v))
		}
	case []uint:
		for _, v := range arr.([]uint) {
			int8Arr = append(int8Arr, int8(v))
		}
	}
	return int8Arr
}
