package application

import (
	"context"
	"now-go-kon/pkg/domain"
)

type AuthService interface {
	Auth(context.Context, domain.UserID, domain.Password) error
}

type UserService interface {
	User(context.Context, domain.UserID) (*domain.User, error)
	UserUpdate(context.Context, *domain.UsersDetails) (*domain.UsersDetails, error)
}
