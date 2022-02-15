package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

// 填充响应信息
func (res *Response) WithMsg(msg string) Response {
	return Response{
		Code: res.Code,
		Msg:  msg,
		Data: res.Data,
	}
}

// 填充响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// 填充响应信息和数据
func (res *Response) WithMsgAndData(msg string, data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  msg,
		Data: data,
	}
}
