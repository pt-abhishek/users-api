package utils

import "net/http"

//RestErr struct
type RestErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

//NewBadRequestError returns a 400
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}
