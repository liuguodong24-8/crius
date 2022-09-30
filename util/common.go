package util

import (
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

// UUIDToString uuid指针转字符串
func UUIDToString(id *uuid.UUID) string {
	if id != nil {
		return id.String()
	}
	return ""
}

// ArrContainElement 数组是否包含
func ArrContainElement(arr, element interface{}) bool {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return false
	}
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == element {
			return true
		}
	}
	return false
}

// TimeUnix32 ...
func TimeUnix32(src *time.Time) int32 {
	if src == nil {
		return 0
	}

	return int32(src.Unix())
}

// TimeUnix64 ...
func TimeUnix64(src *time.Time) int64 {
	if src == nil {
		return 0
	}

	return src.Unix()
}
