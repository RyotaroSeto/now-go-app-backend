package infrastructure

import (
	"context"
	"errors"
	"log"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Board struct {
	ID           int          `gorm:"column:id;primaryKey,omitempty"`
	UserID       int          `gorm:"column:user_id"`
	Body         string       `gorm:"column:body"`
	CreatedDate  time.Time    `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate  time.Time    `gorm:"column:updated_date;autoUpdateTime"`
	UsersDetails UsersDetails `gorm:"foreignKey:UserID;references:UserID"`
}

func (u *Board) toEntity() *domain.Board {
	board := &domain.Board{
		UserID: domain.UserID(u.UserID),
		Body:   domain.Body(u.Body),
	}

	return board
}

func (us *Board) bindEntity(e *domain.Board) {
	u := us.toEntity()
	e.UserID = u.UserID
	e.Body = u.Body
}

func (u *Board) fromEntity(e *domain.Board) {
	u.UserID = e.UserID.Num()
	u.Body = e.Body.String()
}

type BoardRepository struct {
	db *DB
}

var _ domain.BoardRepository = new(BoardRepository)

func NewBoardRepository() *BoardRepository {
	return &BoardRepository{GetDB()}
}

func (u *BoardRepository) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return u.db.Session(&gorm.Session{})
}

func (u *BoardRepository) GetBoard(ctx context.Context, gender domain.Gender) ([]*domain.Board, error) {
	b := []Board{}
	res := u.conn(ctx).Joins("inner join users_details on boards.user_id = users_details.user_id").Where("users_details.gender != ?", gender.String()).Order("boards.created_date desc").Limit(20).Find(&b)
	if res.RowsAffected == 0 {
		msg := "board: is not found"
		return nil, errors.New(msg)
	}
	if err := res.Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}
	log.Println(b)
	// bs := domain.Board{}
	// for _, a := range b {
	// 	bs = append(bs, *a.toEntity())
	// }
	// 	return mls, nil
	return nil, nil
}

func (u *BoardRepository) CreateBoard(ctx context.Context, uParam *domain.Board) (*domain.Board, error) {
	var b Board
	b.fromEntity(uParam)

	if err := u.conn(ctx).Create(&b).Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	return b.toEntity(), nil
}

func (u *BoardRepository) DeleteBoard(ctx context.Context, bID domain.BoardID) error {
	var b Board
	if err := u.conn(ctx).Where(&Board{ID: bID.Num()}).Delete(&b).Error; err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}

	return nil
}
