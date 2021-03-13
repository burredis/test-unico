package app

//Response data model
type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Response normalized error return
func ResponseError(message string, code int) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}

//Response normalized success return
func ResponseSuccess(data interface{}) Response {
	return Response{
		Data: data,
	}
}
