package infrastructure

import (
	"now-go-kon/pkg/domain"
	"time"
)

type UsersDetails struct {
	ID          int       `gorm:"column:id;primaryKey,omitempty"`
	UserID      int       `gorm:"column:user_id"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
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
		ID:          u.ID,
		UserID:      u.UserID,
		DateOfBirth: u.DateOfBirth,
		Gender:      u.Gender,
		Residence:   u.Residence,
		Occupation:  u.Occupation,
		Height:      u.Height,
		Weight:      u.Weight,
	}

	return userDetai
}
