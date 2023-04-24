package infrastructure

import (
	"context"
	"errors"
	"log"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Like struct {
	UserID      int       `gorm:"column:user_id"`
	LikedUserID int       `gorm:"column:liked_user_id"`
	LikedDate   time.Time `gorm:"column:liked_date;autoCreateTime"`
	Status      int       `gorm:"default:1"`
	MessageBody string    `gorm:"column:message_body"`
}

func (u *Like) fromEntity(e *domain.Like) {
	u.UserID = e.UserID.Num()
	u.LikedUserID = e.LikedUserID.Num()
	u.MessageBody = e.MessageBody.String()
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
	var l Like
	l.fromEntity(uParam)

	if err := u.conn(ctx).Create(&l).Error; err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}

	return nil
}
