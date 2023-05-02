package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"log"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID             int          `gorm:"column:id;primaryKey,omitempty"`
	UserName       string       `gorm:"column:user_name"`
	HashedPassword string       `gorm:"column:hashed_password"`
	Status         int          `gorm:"column:status"`
	Email          string       `gorm:"column:email"`
	CreatedDate    time.Time    `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate    time.Time    `gorm:"column:updated_date;autoUpdateTime"`
	UsersDetails   UsersDetails `gorm:"foreignKey:ID;references:UserID"`
}

func (u *Users) toEntity() *domain.User {
	users := &domain.User{
		ID:             domain.UserID(u.ID),
		UserName:       domain.UserName(u.UserName),
		HashedPassword: domain.HashedPassword(u.HashedPassword),
		Status:         domain.Status(u.Status),
		Email:          domain.Email(u.Email),
		UsersDetails:   *u.UsersDetails.toEntity(),
	}

	return users
}

func (us *Users) bindEntity(e *domain.User) {
	u := us.toEntity()
	e.ID = u.ID
	e.UserName = u.UserName
	e.Status = u.Status
	e.Email = u.Email
	e.UsersDetails = u.UsersDetails
}

func (u *Users) fromEntity(e *domain.User) {
	u.ID = e.ID.Num()
	u.UserName = e.UserName.String()
	u.HashedPassword = e.HashedPassword.String()
	u.Status = e.Status.Num()
	u.Email = e.Email.String()
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

func (u *UserRepository) UserCreate(ctx context.Context, uParam *domain.User) (*domain.User, error) {
	us := Users{}
	us.fromEntity(uParam)

	if err := u.conn(ctx).Create(&us).Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	return us.toEntity(), nil
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

	return us.toEntity(), nil
}
