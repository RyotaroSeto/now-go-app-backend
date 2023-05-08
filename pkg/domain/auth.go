package domain

import (
	"time"
)

type Session struct {
	SessionID    string
	UserName     string
	RefreshToken string
	UserAgent    string
	ClientIP     string
	IsBlocked    bool
	ExpiresDate  time.Time
	CreateDate   time.Time
}

type SessionID string

func (s SessionID) String() string {
	return string(s)
}

// func NewSessionID() (SessionID, error) {
// 	u, err := uuid.NewRandom()
// 	if err != nil {
// 		return "", InheritError(InternalServerError, err)
// 	}

// 	return SessionID(u.String()), nil
// }

type RefreshToken string

func (r RefreshToken) String() string {
	return string(r)
}

type UserAgent string

func (u UserAgent) String() string {
	return string(u)
}

type ClientIP string

func (c ClientIP) String() string {
	return string(c)
}
