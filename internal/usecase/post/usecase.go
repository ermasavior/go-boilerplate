package post

import (
	"context"
	"errors"
	"go-boilerplate/internal/model"
)

//go:generate mockgen -package=post -source=usecase.go -destination=mock_usecase.go
type PostRepository interface {
	GetPosts(context.Context) ([]model.Post, error)
	CreatePost(context.Context, model.Post) error
}

type PostUsecase struct {
	repository PostRepository
}

func New(repo PostRepository) *PostUsecase {
	return &PostUsecase{
		repository: repo,
	}
}

func (uc *PostUsecase) GetPosts(ctx context.Context) ([]model.Post, error) {
	return uc.repository.GetPosts(ctx)
}

func (uc *PostUsecase) CreatePost(ctx context.Context, req model.Post) error {
	if req.Title == "" || req.Content == "" {
		return errors.New("invalid params")
	}

	return uc.repository.CreatePost(ctx, req)
}
