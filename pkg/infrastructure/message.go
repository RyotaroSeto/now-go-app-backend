package infrastructure

import (
	"context"
	"now-go-kon/pkg/domain"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *DB
}

var _ domain.MessageRepository = new(MessageRepository)

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{GetDB()}
}

func (u *MessageRepository) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return u.db.Session(&gorm.Session{})
}
