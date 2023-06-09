package domain

import "context"

type AuthRepository interface {
	PasswordAuth(context.Context, UserID, Password) error
	UserGet(context.Context, Email) (*User, error)
	SessionCreate(context.Context, *Session) (*Session, error)
	SessionDelete(context.Context, UserName) error
}

type UserRepository interface {
	UserCreate(context.Context, *User) (*User, error)
	GetUserDetails(context.Context, Gender) ([]*UsersDetails, error)
	GetProfile(context.Context, UserID) (*User, error)
	UpsertProfile(context.Context, *UsersDetails) (*UsersDetails, error)
}

type BoardRepository interface {
	GetScrollBoard(context.Context, Gender, BoardID) ([]*Board, error)
	CreateBoard(context.Context, *Board) (*Board, error)
	DeleteBoard(context.Context, BoardID) error
}

type LikeRepository interface {
	GetLiked(context.Context, UserID) ([]*Like, error)
	CreateLike(context.Context, *Like) error
	ApprovalUser(context.Context, *Like) error
}

type MessageRepository interface {
}
