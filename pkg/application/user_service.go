package application

import (
	"context"
	"fmt"
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

func (s *userService) CreateUser(ctx context.Context, uParam *domain.User) (user *domain.User, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		u, err := s.repo.UserCreate(ctx, uParam)
		if err != nil {
			log.Println(err)
			return err
		}
		user = u
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return user, nil
}

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
		return nil, fmt.Errorf("error: %v", err)
	}

	return user, nil
}

func (s *userService) UserUpsert(ctx context.Context, uParam *domain.UsersDetails) (user *domain.UsersDetails, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		ud, err := s.repo.UpsertProfile(ctx, uParam)
		if err != nil {
			log.Println(err)
			return err
		}
		user = ud
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return user, nil
}
