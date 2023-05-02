package domain

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Status int `json:"status"`
}

func NewErrResponse(status int) *errorResponse {
	return &errorResponse{
		Status: status,
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
