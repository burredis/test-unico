package app

//Response data model
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Response normalized error return
func ResponseMessage(message string) Response {
	return Response{
		Message: message,
	}
}

//Response normalized data return
func ResponseData(data interface{}) Response {
	return Response{
		Data: data,
	}
}
