package models

// APIError define API errors
type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}
