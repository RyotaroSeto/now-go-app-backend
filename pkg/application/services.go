package application

import (
	"context"
	"now-go-kon/pkg/domain"
)

type AuthService interface {
	Auth(context.Context, domain.UserID, domain.Password) error
}
