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

type BoardService interface {
	BoardGet(context.Context, domain.Gender) ([]*domain.Board, error)
	ScrollBoardGet(context.Context, domain.Gender, domain.BoardID) ([]*domain.Board, error)
	BoardCreate(context.Context, *domain.Board) (*domain.Board, error)
	BoardDelete(context.Context, domain.BoardID) error
}

type LikeService interface {
	LikeCreate(context.Context, *domain.Like) error
}

type MessageService interface {
}
