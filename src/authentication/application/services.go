package application

import (
	"authentication/domain"
	"context"
)

type AuthService interface {
	Auth(context.Context, domain.UserID, domain.Password) error
}
