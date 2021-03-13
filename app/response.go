package app

//Response data model
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Response normalized error return
func ResponseError(message string) Response {
	return Response{
		Message: message,
	}
}

//Response normalized success return
func ResponseSuccess(data interface{}) Response {
	return Response{
		Data: data,
	}
}
