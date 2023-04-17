package domain

import "context"

type AuthRepository interface {
	PasswordAuth(context.Context, UserID, Password) error
}

type UserRepository interface {
	GetProfile(context.Context, UserID) (*User, error)
}
