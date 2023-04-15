package domain

import "context"

type AuthRepository interface {
	PasswordAuth(context.Context, UserID, Password) error
}
