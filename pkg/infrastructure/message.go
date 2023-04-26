package infrastructure

import (
	"context"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID             int       `gorm:"column:id;primaryKey,omitempty"`
	SenderUserID   int       `gorm:"column:sender_user_id"`
	ReceiverUserID int       `gorm:"column:receiver_user_id"`
	MessageBody    string    `gorm:"column:message_body"`
	Status         int       `gorm:"default:1"`
	SentDate       time.Time `gorm:"column:sent_date"`
}

func (u *Message) fromEntity(e *domain.Message) {
	u.SenderUserID = e.SenderUserID.Num()
	u.ReceiverUserID = e.ReceiverUserID.Num()
	u.MessageBody = e.MessageBody.String()
	u.SentDate = e.SentDate
}

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
