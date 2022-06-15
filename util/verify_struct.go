package util

import (
	"fmt"

	"github.com/go-playground/validator"
)

// 校验参数
func VerifyParm(s interface{}) error {
	if err := validator.New().Struct(s); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}
		return err
	}
	return nil
}
