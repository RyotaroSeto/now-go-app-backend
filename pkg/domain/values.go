package domain

import "time"

type User struct {
	ID           UserID
	UserName     UserName
	Status       Status
	Email        Email
	UsersDetails UsersDetails
}

func NewUser(uID UserID, un UserName, s Status, e Email) *User {
	return &User{
		ID:       uID,
		UserName: un,
		Status:   s,
		Email:    e,
	}
}

type UserID int

func (u UserID) Num() int {
	return int(u)
}

type UserName string

func (u UserName) String() string {
	return string(u)
}

type Status int

func (s Status) Num() int {
	return int(s)
}

type Email string

func (e Email) String() string {
	return string(e)
}

type UsersDetails struct {
	ID          int
	UserID      int
	DateOfBirth time.Time
	Gender      string
	Residence   string
	Occupation  string
	Height      int
	Weight      int
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Password string

func (p Password) String() string {
	return string(p)
}
