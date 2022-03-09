package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyResponse struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, error string, data interface{}) Response {
	return Response{
		Status:  false,
		Message: message,
		Data:    data,
		Error:   strings.Split(error, "\n"),
	}
}
