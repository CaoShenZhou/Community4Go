package util

import (
	"math/rand"
	"strings"
)

// 获取随机字符串
func RandomString(n int, allowedChars string) string {
	var defChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var chars []byte
	if len(strings.Trim(allowedChars, "")) == 0 {
		chars = []byte(defChars)
	} else {
		chars = []byte(chars)
	}
	bytes := []byte(chars)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
