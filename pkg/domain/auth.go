package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	SessionID    SessionID
	UserName     UserName
	RefreshToken RefreshToken
	UserAgent    UserAgent
	ClientIP     ClientIP
	IsBlocked    IsBlocked
	ExpiresDate  time.Time
	CreateDate   time.Time
}

type SessionID uuid.UUID

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

type IsBlocked bool
