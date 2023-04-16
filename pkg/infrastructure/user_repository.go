package infrastructure

import (
	"context"
	"now-go-kon/pkg/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *DB
}

var _ domain.UserRepository = new(UserRepository)

func NewUserRepository() *UserRepository {
	return &UserRepository{GetDB()}
}

func (u *UserRepository) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return u.db.Session(&gorm.Session{})
}

func (u *UserRepository) GetProfile(ctx context.Context, uID domain.UserID) error {
	return nil
}
