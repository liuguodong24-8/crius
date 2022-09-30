package util

import (
	"fmt"
	"time"
)

var staffCodePre = []byte{}

// GenerateStaffCode 生成员工编号
func GenerateStaffCode(num int64) string {
	prefix := 97 + (time.Now().Year()-2017)%26
	return fmt.Sprintf("%c%06d", prefix, num)
}

// GenerateBranchCode 生成门店编号
func GenerateBranchCode(provinceID string, num int64) string {
	return fmt.Sprintf("%s%06d", provinceID, num)
}
