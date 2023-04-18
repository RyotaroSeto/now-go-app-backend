package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"log"
	"now-go-kon/pkg/domain"

	"gorm.io/gorm"
)

type Users struct {
	ID           int          `gorm:"column:id;primaryKey,omitempty"`
	UserName     string       `gorm:"column:user_name"`
	Password     string       `gorm:"column:password"`
	Email        string       `gorm:"column:email"`
	UsersDetails UsersDetails `gorm:"foreignKey:ID;references:UserID"`
}

func (u *Users) toEntity() *domain.User {
	users := &domain.User{
		ID:       domain.UserID(u.ID),
		UserName: domain.UserName(u.UserName),
		Password: domain.Password(u.Password),
		Email:    domain.Email(u.Email),
	}

	return users
}

func (us *Users) bindEntity(e *domain.User) {
	u := us.toEntity()
	e.ID = u.ID
	e.UserName = u.UserName
	e.Password = u.Password
	e.Email = u.Email
}

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

func (u *UserRepository) GetProfile(ctx context.Context, uID domain.UserID) (*domain.User, error) {
	us := Users{}
	q := Users{ID: uID.Num()}
	res := u.conn(ctx).Preload("UsersDetails").Where(&q).First(&us)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := fmt.Sprintf("UserID: %d is not found", uID.Num())
			return nil, errors.New(msg)

		}
		return nil, err
	}

	log.Println(us)
	return us.toEntity(), nil
}
