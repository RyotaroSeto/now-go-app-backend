package application

import (
	"context"
	"log"
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

func (s *userService) User(ctx context.Context, uID domain.UserID) (user *domain.User, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		u, err := s.repo.GetProfile(ctx, uID)
		if err != nil {
			log.Println(err)
			return err
		}
		user = u
		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, err
}
