package serializer

// Response 统一响应结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data interface{} `json:"data"`
}
