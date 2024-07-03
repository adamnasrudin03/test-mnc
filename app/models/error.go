package models

type RespError struct {
	StatusCode int   `json:"status_code"`
	Error      error `json:"error"`
}
