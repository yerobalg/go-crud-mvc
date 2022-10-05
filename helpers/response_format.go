package helper

type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func (h *Helper) ResponseFormat(message string, success bool, data interface{}) Response {
	return Response{
		Message: message,
		Success: success,
		Data:    data,
	}
}
