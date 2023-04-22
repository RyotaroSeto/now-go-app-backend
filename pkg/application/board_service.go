package application

import (
	"context"
	"log"
	"now-go-kon/pkg/domain"
)

type boardService struct {
	repo domain.BoardRepository
	tx   Transaction
}

func NewBoardService(repo domain.BoardRepository, tx Transaction) BoardService {
	return &boardService{
		repo: repo,
		tx:   tx,
	}
}

var _ BoardService = new(boardService)

func (s *boardService) BoardGet(ctx context.Context, gender domain.Gender) (boards []*domain.Board, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		b, err := s.repo.GetBoard(ctx, gender)
		log.Println(b)
		if err != nil {
			log.Println(err)
			return err
		}
		// board = b
		return nil
	})
	return nil, nil
	// if err != nil {
	// 	return nil, err
	// }

	// return board, err
}

func (s *boardService) BoardCreate(ctx context.Context, uParam *domain.Board) (board *domain.Board, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		b, err := s.repo.CreateBoard(ctx, uParam)
		if err != nil {
			log.Println(err)
			return err
		}
		board = b
		return nil
	})
	if err != nil {
		return nil, err
	}

	return board, err
}

func (s *boardService) BoardDelete(ctx context.Context, bID domain.BoardID) (board *domain.Board, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		b, err := s.repo.DeleteBoard(ctx, bID)
		if err != nil {
			log.Println(err)
			return err
		}
		board = b
		return nil
	})
	if err != nil {
		return nil, err
	}

	return board, err
}
