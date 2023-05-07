package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"now-go-kon/pkg/domain"

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
func (r *AuthRepository) UserGet(ctx context.Context, email domain.Email) (*domain.User, error) {
	us := Users{}
	q := Users{Email: email.String()}
	res := r.conn(ctx).Where(&q).First(&us)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := fmt.Sprintf("Email: %s is not found", email.String())
			return nil, errors.New(msg)
		}
		return nil, err
	}

	return us.toEntity(), nil
}

func (r *AuthRepository) SessionCreate(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	return nil, nil
}

func (r *AuthRepository) PasswordAuth(ctx context.Context, uID domain.UserID, password domain.Password) error {
	// パスワード認証の処理は割愛
	return nil
}
