package application

import (
	"context"
	"log"
	"now-go-kon/pkg/domain"
)

type likeService struct {
	repo domain.LikeRepository
	tx   Transaction
}

func NewLikeService(repo domain.LikeRepository, tx Transaction) LikeService {
	return &likeService{
		repo: repo,
		tx:   tx,
	}
}

var _ LikeService = new(likeService)

func (s *likeService) LikeCreate(ctx context.Context, uParam *domain.Like) (err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := s.repo.CreateLike(ctx, uParam); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return err
}

func (s *likeService) LikeGet(ctx context.Context, uID domain.UserID) (likes []*domain.Like, err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		l, err := s.repo.GetLiked(ctx, uID)
		if err != nil {
			log.Println(err)
			return err
		}
		likes = l
		return nil
	})
	if err != nil {
		return nil, err
	}

	return likes, err
}

func (s *likeService) Approval(ctx context.Context, uParam *domain.Like) (err error) {
	err = s.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := s.repo.ApprovalUser(ctx, uParam); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return err
}
