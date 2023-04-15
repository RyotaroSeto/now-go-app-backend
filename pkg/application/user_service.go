package application

import (
	"context"
	"now-go-kon/pkg/domain"
)

type userService struct {
	repo domain.UserRepository
	tx   Transaction
}

func NewUserService(repo domain.UserRepository, tx Transaction) UserService {
	return &userService{
		repo: repo,
		tx:   tx,
	}
}

var _ UserService = new(userService)

func (s *userService) User(ctx context.Context, uID domain.UserID) error {
	err := s.tx.Transaction(ctx, func(ctx context.Context) error {
		return s.repo.GetProfile(ctx, uID)
	})

	return err
}
