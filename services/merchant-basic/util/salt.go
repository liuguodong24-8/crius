package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// DefaultPassword 员工默认密码
const DefaultPassword = "123456"

var saltByteArr = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

// GenerateSalt 根据当前时间纳秒生成一个随机13位加密salt
func GenerateSalt() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var buff bytes.Buffer
	for i := 0; i < 13; i++ {
		n := r.Intn(36)
		buff.WriteByte(saltByteArr[n])
	}
	return buff.String()
}

// GetMD5 生成输入字符串的md5值
func GetMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
