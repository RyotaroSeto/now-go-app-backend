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

func (u *Like) toEntity() *domain.Like {
	board := &domain.Like{
		UserID:      domain.UserID(u.UserID),
		LikedUserID: domain.UserID(u.LikedUserID),
		LikedDate:   u.LikedDate,
		Status:      domain.Status(u.Status),
		MessageBody: domain.MessageBody(u.MessageBody),
	}

	return board
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

func (u *LikeRepository) GetLiked(ctx context.Context, uID domain.UserID) ([]*domain.Like, error) {
	l := []Like{}

	q := Like{UserID: uID.Num(), Status: 1}
	res := u.conn(ctx).Where(&q).Order("likes.liked_date desc").Find(&l)
	if res.RowsAffected == 0 {
		msg := "user_id: is not found"
		return nil, errors.New(msg)
	}
	if err := res.Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	ls := []*domain.Like{}
	for _, a := range l {
		ls = append(ls, a.toEntity())
	}
	return ls, nil
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

func (u *LikeRepository) ApprovalUser(ctx context.Context, uParam *domain.Like) error {
	q := Like{UserID: uParam.UserID.Num(), LikedUserID: uParam.LikedUserID.Num()}
	if err := u.conn(ctx).Model(&Like{}).Where(&q).Update("status", 0).Error; err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}
	var l Like
	res := u.conn(ctx).Where(&q).First(&l)
	if res.RowsAffected == 0 {
		msg := "user_id: is not found"
		return errors.New(msg)
	}
	if err := res.Error; err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}

	var m Message
	mParam := &domain.Message{
		SenderUserID:   uParam.UserID,
		ReceiverUserID: uParam.LikedUserID,
		MessageBody:    domain.MessageBody(l.MessageBody),
		SentDate:       l.LikedDate,
	}
	m.fromEntity(mParam)
	if err := u.conn(ctx).Create(&m).Error; err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}

	return nil
}
