package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyResponse struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, errors, data interface{}) Response {
	switch v := errors.(type) {
	case map[string]interface{}:
		errors = v
	case string:
		errors = strings.Split(v, "\n")
	}

	return Response{
		Status:  false,
		Message: message,
		Data:    data,
		Errors:  errors,
	}
}
