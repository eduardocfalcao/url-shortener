package model

type ErrorResponse struct {
	Error     string         `json:"error"`
	ErrorCode ErrorCode      `json:"error_code"`
	Details   []ErrorDetails `json:"details,omitempty"`
}

type ErrorDetails struct {
	Message string `json:"message"`
}

type ErrorCode string

const (
	InvalidRequest = ErrorCode("Invalid Request")
)
