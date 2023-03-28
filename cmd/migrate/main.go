package main

import (
	"go-boilerplate/pkg/config"
	db "go-boilerplate/pkg/database"
	"log"
)

func main() {
	config.Load()

	if err := db.RunLatestMigration(); err != nil {
		log.Fatalln("error migrate db", err)
	}
}
