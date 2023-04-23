package infrastructure

import (
	"context"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Like struct {
	UserID      int       `gorm:"column:user_id"`
	LikedUserID int       `gorm:"column:liked_user_id"`
	LikedDate   time.Time `gorm:"column:liked_date;autoCreateTime"`
	Status      int       `gorm:"column:status"`
	MessageBody string    `gorm:"column:message_body"`
}

type LikeRepository struct {
	db *DB
}

var _ domain.LikeRepository = new(LikeRepository)

func NewLikeRepository() *LikeRepository {
	return &LikeRepository{GetDB()}
}

func (u *LikeRepository) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return u.db.Session(&gorm.Session{})
}

func (u *LikeRepository) CreateLike(ctx context.Context, uParam *domain.Like) error {
	return nil
}
