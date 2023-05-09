package application

import (
	"context"
	"fmt"
	"log"
	"now-go-kon/pkg/domain"
)

type authService struct {
	repo domain.AuthRepository
	tx   Transaction
}

func NewAuthService(repo domain.AuthRepository, tx Transaction) AuthService {
	return &authService{
		repo: repo,
		tx:   tx,
	}
}

var _ AuthService = new(authService)

func (s *authService) GetUser(ctx context.Context, email domain.Email) (user *domain.User, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		l, err := s.repo.UserGet(ctx, email)
		if err != nil {
			log.Println(err)
			return err
		}
		user = l
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return user, nil
}

func (s *authService) CreateSession(ctx context.Context, sParam *domain.Session) (session *domain.Session, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		err := s.repo.SessionDelete(ctx, domain.UserName(sParam.UserName))
		if err != nil {
			log.Println(err)
			return err
		}

		l, err := s.repo.SessionCreate(ctx, sParam)
		if err != nil {
			log.Println(err)
			return err
		}
		session = l
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return session, nil
}
