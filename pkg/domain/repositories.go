package domain

import "context"

type AuthRepository interface {
	PasswordAuth(context.Context, UserID, Password) error
}

type UserRepository interface {
	GetProfile(context.Context, UserID) (*User, error)
	UpdateProfile(context.Context, *UsersDetails) (*UsersDetails, error)
}

type BoardRepository interface {
	CreateBoard(context.Context, *Board) (*Board, error)
}
