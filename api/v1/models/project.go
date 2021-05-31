package models

type AddProjectRequest struct {
	ContactEmail string `json:"contactEmail"`
}

type AddProjectResponse struct {
	BaseResponse
	ApiKey string `json:"apiKey"`
}

// Validate model
func (req *AddProjectRequest) IsValid() bool {
	return len(req.ContactEmail) > 0
}
