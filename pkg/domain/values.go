package domain

import "time"

type User struct {
	ID             UserID
	UserName       UserName
	HashedPassword HashedPassword
	Status         Status
	Email          Email
	CreateDate     time.Time
	UpdatedDate    time.Time
	UsersDetails   UsersDetails
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

type HashedPassword string

func (h HashedPassword) String() string {
	return string(h)
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
	ID           int
	UserID       UserID
	Name         string
	Age          int
	Gender       string
	Height       int
	Location     string
	Work         string
	Graduation   string
	Hobby        string
	Passion      string
	Tweet        string
	Introduction string
	CreatedDate  time.Time
	UpdatedDate  time.Time
}

type UserDetailID int

func (u UserDetailID) Num() int {
	return int(u)
}

type DateOfBirth int

func (d DateOfBirth) Num() int {
	return int(d)
}

type Gender string

func (g Gender) String() string {
	return string(g)
}

type Residence string

func (r Residence) String() string {
	return string(r)
}

type Occupation string

func (o Occupation) String() string {
	return string(o)
}

type Height int

func (h Height) Num() int {
	return int(h)
}

type Weight int

func (w Weight) Num() int {
	return int(w)
}

type Password string

func (p Password) String() string {
	return string(p)
}
