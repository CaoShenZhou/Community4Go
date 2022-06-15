package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 生成MD5
func generateMd5(str string) string {
	data := []byte(str)
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	return s
}
