package domain

import "time"

type Message struct {
	ID             MessageID
	SenderUserID   UserID
	ReceiverUserID UserID
	MessageBody    MessageBody
	SentDate       time.Time
}

type MessageID int

func (m MessageID) Num() int {
	return int(m)
}
