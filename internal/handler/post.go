package handler

import (
	"context"
	"net/http"

	"go-boilerplate/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -package=handler -source=post.go -destination=mock/mock_post.go
type PostUsecase interface {
	GetPosts(ctx context.Context) ([]model.Post, error)
	CreatePost(ctx context.Context, req model.Post) error
}

type PostHandler struct {
	usecase PostUsecase
}

func NewPostHandler(uc PostUsecase) PostHandler {
	return PostHandler{
		usecase: uc,
	}
}

func (h *PostHandler) Hello(c *gin.Context) {
	post := model.Post{
		Title:   "hello",
		Content: "hello world",
	}

	c.IndentedJSON(http.StatusOK, newResponse(post))
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.usecase.GetPosts(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, newResponse(posts))

}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req model.Post
	if err := getRequestBody(c, &req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newErrorResponse(errInvalidBodyRequest))
		return
	}

	if err := h.usecase.CreatePost(c, req); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusCreated, newResponse("success"))
}
