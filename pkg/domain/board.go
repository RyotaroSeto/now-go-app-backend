package domain

import "time"

type Board struct {
	ID           BoardID
	UserID       UserID
	Body         Body
	CreatedDate  time.Time
	UsersDetails UsersDetails
}

type BoardID int

func (b BoardID) Num() int {
	return int(b)
}

type Body string

func (b Body) String() string {
	return string(b)
}
