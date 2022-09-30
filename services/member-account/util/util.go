package util

import (
	"crypto/sha256"
	"encoding/hex"
)

var h = sha256.New()

// DefaultCardPassword 卡默认密码
var DefaultCardPassword = Sha256("888888")

// Sha256 sha256哈希
func Sha256(s string) string {
	h.Write([]byte(s))
	b := h.Sum(nil)
	h.Reset()
	return hex.EncodeToString(b)
}
