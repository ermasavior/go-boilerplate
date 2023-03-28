package handler

import "github.com/gin-gonic/gin"

type Handlers struct {
	Post PostHandler
}

func NewRouter(h *Handlers) *gin.Engine {
	r := gin.Default()

	r.GET("/", h.Post.Hello)

	v1 := r.Group("/posts")
	{
		v1.GET("", h.Post.GetPosts)
		v1.POST("", h.Post.CreatePost)
	}

	return r
}
