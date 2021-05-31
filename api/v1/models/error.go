package models

type ErrorResponse struct {
	BaseResponse
	Message string `json:"message"`
}
