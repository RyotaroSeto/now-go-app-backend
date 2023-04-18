package domain

type ErrorResponse struct {
	Status int `json:"status"`
}

func NewErrResponse(status int) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
	}
}
