package application

import "now-go-kon/pkg/domain"

type messageService struct {
	repo domain.MessageRepository
	tx   Transaction
}

func NewMessageService(repo domain.MessageRepository, tx Transaction) MessageService {
	return &messageService{
		repo: repo,
		tx:   tx,
	}
}

var _ MessageService = new(messageService)
