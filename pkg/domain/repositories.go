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
	GetBoard(context.Context, Gender) ([]*Board, error)
	GetScrollBoard(context.Context, Gender, BoardID) ([]*Board, error)
	CreateBoard(context.Context, *Board) (*Board, error)
	DeleteBoard(context.Context, BoardID) error
}

type LikeRepository interface {
	CreateLike(context.Context, *Like) error
}

type MessageRepository interface {
}
