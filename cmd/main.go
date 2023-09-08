package main

import (
	"github.com/spf13/viper"
	todo "github.com/tuanneli/REST-api.git"
	"github.com/tuanneli/REST-api.git/pkg/handler"
	"github.com/tuanneli/REST-api.git/pkg/repository"
	"github.com/tuanneli/REST-api.git/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())

	if err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigFile("configs")
	viper.SetConfigFile("config")
	viper.SetConfigType("json")
	return viper.ReadInConfig()
}
