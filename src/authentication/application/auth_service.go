package application

import (
	"authentication/domain"
	"context"
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

func (s *authService) Auth(ctx context.Context, uID domain.UserID, password domain.Password) error {
	err := s.tx.Transaction(ctx, func(ctx context.Context) error {
		return s.repo.PasswordAuth(ctx, uID, password)
	})

	return err
}
