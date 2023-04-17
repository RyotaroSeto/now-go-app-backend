package infrastructure

import "time"

type UserDetail struct {
	// ID          int       `gorm:"column:id;primaryKey,omitempty"`
	UserID      int       `gorm:"column:user_id"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
	Gender      string    `gorm:"column:gender"`
	Residence   string    `gorm:"column:residence"`
	Occupation  string    `gorm:"column:occupation"`
	Height      int       `gorm:"column:height"`
	Weight      int       `gorm:"column:weight"`
}
