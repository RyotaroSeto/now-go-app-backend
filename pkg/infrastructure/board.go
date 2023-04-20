package infrastructure

import (
	"context"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Board struct {
	ID          int       `gorm:"column:id;primaryKey,omitempty"`
	UserID      int       `gorm:"column:user_id"`
	Body        string    `gorm:"column:body"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime"`
}

// func (u *Users) toEntity() *domain.User {
// 	users := &domain.User{
// 		ID:           domain.UserID(u.ID),
// 		UserName:     domain.UserName(u.UserName),
// 		Status:       domain.Status(u.Status),
// 		Email:        domain.Email(u.Email),
// 		UsersDetails: *u.UsersDetails.toEntity(),
// 	}

// 	return users
// }

// func (us *Users) bindEntity(e *domain.User) {
// 	u := us.toEntity()
// 	e.ID = u.ID
// 	e.UserName = u.UserName
// 	e.Status = u.Status
// 	e.Email = u.Email
// 	e.UsersDetails = u.UsersDetails
// }

// func (u *Users) fromEntity(e *domain.User) {
// 	u.ID = e.ID.Num()
// 	u.UserName = e.UserName.String()
// 	u.Status = e.Status.Num()
// 	u.Email = e.Email.String()
// }

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

func (u *BoardRepository) CreateBoard(ctx context.Context, uID domain.UserID) (*domain.Board, error) {
	// us := Users{}
	// q := Users{ID: uID.Num()}
	// res := u.conn(ctx).Preload("UsersDetails").Where(&q).First(&us)
	// if err := res.Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		msg := fmt.Sprintf("UserID: %d is not found", uID.Num())
	// 		return nil, errors.New(msg)

	// 	}
	// 	return nil, err
	// }

	return nil, nil
}
