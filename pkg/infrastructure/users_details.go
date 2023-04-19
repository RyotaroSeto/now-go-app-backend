package infrastructure

import (
	"context"
	"now-go-kon/pkg/domain"
	"time"
)

type UsersDetails struct {
	ID          int       `gorm:"column:id;primaryKey,omitempty"`
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
		ID:          domain.UserDetailID(u.ID),
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

func (u *UserRepository) UpdateProfile(ctx context.Context, uParam domain.UsersDetails) (*domain.UsersDetails, error) {
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
	// return us.toEntity(), nil
}
