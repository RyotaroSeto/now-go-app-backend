package infrastructure

import (
	"authentication/domain"
	"context"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *DB
}

var _ domain.AuthRepository = new(AuthRepository)

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{GetDB()}
}

func (r *AuthRepository) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return r.db.Session(&gorm.Session{})
}

func (r *AuthRepository) PasswordAuth(ctx context.Context, uID domain.UserID, password domain.Password) error {
	// パスワード認証の処理は割愛
	return nil
}
