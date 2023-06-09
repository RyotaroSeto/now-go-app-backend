package infrastructure

import (
	"context"
	"errors"
	"log"
	"now-go-kon/pkg/domain"
	"time"
)

type UsersDetails struct {
	UserID       int       `gorm:"column:user_id"`
	Name         string    `gorm:"column:name"`
	Age          int       `gorm:"column:age"`
	Gender       string    `gorm:"column:gender"`
	Height       int       `gorm:"column:height"`
	Location     string    `gorm:"column:location"`
	Work         string    `gorm:"column:work"`
	Graduation   string    `gorm:"column:graduation"`
	Hobby        string    `gorm:"column:hobby"`
	Passion      string    `gorm:"column:passion"`
	Tweet        string    `gorm:"column:tweet"`
	Introduction string    `gorm:"column:introduction"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime"`
	Board        Board     `gorm:"foreignKey:UserID;references:UserID"`
}

func (u *UsersDetails) toEntity() *domain.UsersDetails {
	userDetai := &domain.UsersDetails{
		UserID:       domain.UserID(u.UserID),
		Name:         u.Name,
		Age:          u.Age,
		Gender:       u.Gender,
		Height:       u.Height,
		Location:     u.Location,
		Work:         u.Work,
		Graduation:   u.Graduation,
		Hobby:        u.Hobby,
		Passion:      u.Passion,
		Tweet:        u.Tweet,
		Introduction: u.Introduction,
		Board:        *u.Board.toEntity(),
	}

	return userDetai
}

func (us *UsersDetails) bindEntity(e *domain.UsersDetails) {
	u := us.toEntity()
	e.UserID = u.UserID
	e.Name = u.Name
	e.Age = u.Age
	e.Gender = u.Gender
	e.Height = u.Height
	e.Location = u.Location
	e.Work = u.Work
	e.Graduation = u.Graduation
	e.Hobby = u.Hobby
	e.Passion = u.Passion
	e.Tweet = u.Tweet
	e.Introduction = u.Introduction
	e.Board = u.Board
}

func (u *UsersDetails) fromEntity(e *domain.UsersDetails) {
	var ud Board
	ud.fromEntity(&e.Board)
	u.UserID = e.UserID.Num()
	u.Name = e.Name
	u.Age = e.Age
	u.Gender = e.Gender
	u.Height = e.Height
	u.Location = e.Location
	u.Work = e.Work
	u.Graduation = e.Graduation
	u.Hobby = e.Hobby
	u.Passion = e.Passion
	u.Tweet = e.Tweet
	u.Introduction = e.Introduction
	u.Board = ud
}

func (u *UserRepository) UpsertProfile(ctx context.Context, uParam *domain.UsersDetails) (*domain.UsersDetails, error) {
	var param, ud UsersDetails
	param.fromEntity(uParam)

	q := UsersDetails{UserID: uParam.UserID.Num()}
	res := u.conn(ctx).Where(&q).Assign(&param).FirstOrCreate(&ud)
	if err := res.Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	return ud.toEntity(), nil
}

func (u *UserRepository) GetUserDetails(ctx context.Context, gender domain.Gender) ([]*domain.UsersDetails, error) {
	ud := []UsersDetails{}

	q := UsersDetails{Gender: gender.String()}
	res := u.conn(ctx).Preload("Board").Where(&q).Limit(20).Find(&ud)
	if res.RowsAffected == 0 {
		msg := "board: is not found"
		return nil, errors.New(msg)
	}
	if err := res.Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	uds := []*domain.UsersDetails{}
	for _, a := range ud {
		uds = append(uds, a.toEntity())
	}

	return uds, nil
}
