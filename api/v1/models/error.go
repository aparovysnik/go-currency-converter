package models

type StatusResponse struct {
	BaseResponse
	Message string `json:"message"`
}
