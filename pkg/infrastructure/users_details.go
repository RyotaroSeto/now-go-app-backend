package infrastructure

import (
	"context"
	"errors"
	"log"
	"now-go-kon/pkg/domain"
	"time"
)

type UsersDetails struct {
	UserID      int       `gorm:"column:user_id"`
	DateOfBirth int       `gorm:"column:date_of_birth"`
	Gender      string    `gorm:"column:gender"`
	Residence   string    `gorm:"column:residence"`
	Occupation  string    `gorm:"column:occupation"`
	Height      int       `gorm:"column:height"`
	Weight      int       `gorm:"column:weight"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime"`
}

func (u *UsersDetails) toEntity() *domain.UsersDetails {
	userDetai := &domain.UsersDetails{
		UserID:      domain.UserID(u.UserID),
		DateOfBirth: domain.DateOfBirth(u.DateOfBirth),
		Gender:      domain.Gender(u.Gender),
		Residence:   domain.Residence(u.Residence),
		Occupation:  domain.Occupation(u.Occupation),
		Height:      domain.Height(u.Height),
		Weight:      domain.Weight(u.Weight),
	}

	return userDetai
}

func (us *UsersDetails) bindEntity(e *domain.UsersDetails) {
	u := us.toEntity()
	e.UserID = u.UserID
	e.DateOfBirth = u.DateOfBirth
	e.Gender = u.Gender
	e.Residence = u.Residence
	e.Occupation = u.Occupation
	e.Height = u.Height
	e.Weight = u.Weight
}

func (u *UsersDetails) fromEntity(e *domain.UsersDetails) {
	u.UserID = e.UserID.Num()
	u.DateOfBirth = e.DateOfBirth.Num()
	u.Gender = e.Gender.String()
	u.Residence = e.Residence.String()
	u.Occupation = e.Occupation.String()
	u.Height = e.Height.Num()
	u.Weight = e.Weight.Num()
}

func (u *UserRepository) UpdateProfile(ctx context.Context, uParam *domain.UsersDetails) (*domain.UsersDetails, error) {
	var ud UsersDetails
	ud.fromEntity(uParam)

	q := UsersDetails{UserID: uParam.UserID.Num()}
	if err := u.conn(ctx).Where(&q).Save(&ud).Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	return ud.toEntity(), nil
}
