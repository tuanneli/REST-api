package main

import (
	todo "github.com/tuanneli/REST-api.git"
	"github.com/tuanneli/REST-api.git/pkg/handler"
	"github.com/tuanneli/REST-api.git/pkg/repository"
	"github.com/tuanneli/REST-api.git/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	err := srv.Run("8000", handlers.InitRoutes())

	if err != nil {
		log.Fatal("An error occurred: %s", err.Error())
	}
}
