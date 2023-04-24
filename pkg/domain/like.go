package domain

import "time"

type Like struct {
	UserID      UserID
	LikedUserID UserID
	LikedDate   time.Time
	Status      Status
	MessageBody MessageBody
}

type MessageBody string

func (m MessageBody) String() string {
	return string(m)
}
