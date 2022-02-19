package response

var (
	Ok = response(0, "成功")
	// 内部问题1开头
	ServerError   = response(1001, "内部服务异常")
	EmailError    = response(1002, "邮箱服务异常")
	RedisError    = response(1003, "缓存服务异常")
	DatabaseError = response(1004, "数据库服务异常")
	// 外部问题2开头
	BadRequest      = response(2001, "请求有误")
	InvalidParams   = response(2002, "参数有误")
	Unauthenticated = response(2003, "未认证")
	Unauthorized    = response(2004, "未授权")
	// 业务问题3开头
	NotFound     = response(3001, "找不到")
	Unavailable  = response(3002, "不可用")
	ValidateFail = response(3003, "验证失败")
	// 三方问题4开头
)
