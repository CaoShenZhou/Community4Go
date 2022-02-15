package response

var (
	Ok              = response(200, "成功")
	BadRequest      = response(400, "坏请求")
	Unauthenticated = response(401, "未认证")
	Unauthorized    = response(403, "未授权")
	NotFound        = response(404, "找不到")
	InvalidParams   = response(412, "入参错误")
	ServerError     = response(500, "服务内部错误")
)
