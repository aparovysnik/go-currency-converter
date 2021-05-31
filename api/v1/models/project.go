package models

type AddProjectRequest struct {
	ContactEmail string
}

type AddProjectResponse struct {
	BaseResponse
	ApiKey string
}

// Validate model
func (req *AddProjectRequest) IsValid() bool {
	return len(req.ContactEmail) > 0
}
