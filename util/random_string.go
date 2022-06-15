package util

import (
	"math/rand"
)

// 获取随机字符串 []rune("你好世界")
func RandomString(n int, allowedChars []rune) string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars
	}
	b := make([]rune, n)
	for i := range b {
		// rand.Seed(time.Now().Unix())
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
