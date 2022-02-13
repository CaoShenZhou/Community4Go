package util

import (
	"strings"

	"github.com/google/uuid"
)

// 获取UUID
func GetUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
