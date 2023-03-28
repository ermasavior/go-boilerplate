package post

import (
	"context"
	"errors"
	"go-boilerplate/internal/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type TestSuite struct {
	mRepo *MockPostRepository
}

func initTestSuite(t *testing.T) *TestSuite {
	ctrl := gomock.NewController(t)
	return &TestSuite{
		mRepo: NewMockPostRepository(ctrl),
	}
}

func TestUsecase_GetPosts(t *testing.T) {
	suite := initTestSuite(t)
	uc := New(suite.mRepo)

	expected := []model.Post{
		{
			Title:   "hello",
			Content: "hello world",
		},
	}

	t.Run("success", func(t *testing.T) {
		suite.mRepo.EXPECT().GetPosts(gomock.Any()).Return(expected, nil)
		got, err := uc.GetPosts(context.Background())

		assert.Equal(t, true, err == nil)
		assert.Equal(t, expected, got)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		suite.mRepo.EXPECT().GetPosts(gomock.Any()).Return(nil, errors.New("error"))
		_, err := uc.GetPosts(context.Background())

		assert.Equal(t, false, err == nil)
	})
}

func TestUsecase_CreatePost(t *testing.T) {
	suite := initTestSuite(t)
	uc := New(suite.mRepo)

	req := model.Post{
		Title:   "hello",
		Content: "hello world",
	}

	t.Run("success", func(t *testing.T) {
		suite.mRepo.EXPECT().CreatePost(gomock.Any(), req).Return(nil)
		err := uc.CreatePost(context.Background(), req)

		assert.Equal(t, true, err == nil)
	})

	t.Run("invalid req", func(t *testing.T) {
		invalidReq := model.Post{
			Title:   "",
			Content: "",
		}
		err := uc.CreatePost(context.Background(), invalidReq)

		assert.Equal(t, false, err == nil)
	})

	t.Run("usecase returns error", func(t *testing.T) {
		suite.mRepo.EXPECT().CreatePost(gomock.Any(), req).Return(errors.New("error"))
		err := uc.CreatePost(context.Background(), req)

		assert.Equal(t, false, err == nil)
	})
}
