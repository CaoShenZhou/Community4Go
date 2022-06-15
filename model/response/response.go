package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0 // 成功
	FAIL    = 1 // 失败
	ERROR   = 2 // 错误
)

const (
	PARM_BIND_FAIL   = "参数绑定失败"
	PARM_VERIFY_FAIL = "参数验证失败"
)

// 业务响应状态
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// ----------成功

func OK(c *gin.Context) {
	Result(c, SUCCESS, "", nil)
}

func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg, nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "", data)
}

func OkWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	Result(c, SUCCESS, msg, data)
}

// ----------失败

func Fail(c *gin.Context) {
	Result(c, FAIL, "", nil)
}

func FailWithMsg(c *gin.Context, msg string) {
	Result(c, FAIL, msg, nil)
}

func FailWithData(c *gin.Context, data interface{}) {
	Result(c, FAIL, "", data)
}

func FailWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	Result(c, FAIL, msg, data)
}

// ----------错误

func Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, http.StatusText(code))
}
