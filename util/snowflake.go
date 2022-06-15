package util

import (
	"github.com/bwmarrin/snowflake"
)

// 获取雪花ID
func GetSnowflake() (*snowflake.ID, error) {
	if node, err := snowflake.NewNode(1); err != nil {
		return nil, err
	} else {
		id := node.Generate()
		return &id, nil
	}
}
