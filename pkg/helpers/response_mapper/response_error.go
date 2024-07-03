package response_mapper

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type ResponseDefault struct {
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func APIResponse(message string, statusCode int, data interface{}) ResponseDefault {
	status := ""
	switch statusCode {
	case http.StatusOK:
		status = "Success"
	default:
		status = http.StatusText(statusCode)
	}

	if message != "" {
		status = ""
	}

	return ResponseDefault{
		Status:  helpers.ToUpper(status),
		Message: helpers.ToTitle(message),
		Result:  data,
	}
}
