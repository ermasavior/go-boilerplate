package main

import (
	"fmt"
	"go-boilerplate/internal/handler"
	postRepository "go-boilerplate/internal/repository/post"
	postUsecase "go-boilerplate/internal/usecase/post"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/database"
	"log"
)

func main() {
	config.Load()

	db, err := database.NewDB()
	if err != nil {
		log.Fatalln("error connecting to db")
	}

	postRepo := postRepository.New(db)
	postUC := postUsecase.New(postRepo)
	postHandler := handler.NewPostHandler(postUC)

	h := &handler.Handlers{
		Post: postHandler,
	}
	r := handler.NewRouter(h)

	port := fmt.Sprintf(":%s", config.Get().AppPort)
	r.Run(port)
}
