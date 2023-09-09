package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	todo "github.com/tuanneli/REST-api.git"
	"github.com/tuanneli/REST-api.git/pkg/handler"
	"github.com/tuanneli/REST-api.git/pkg/repository"
	"github.com/tuanneli/REST-api.git/pkg/service"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("faild initializing db: %s", err.Error())
	}

	port := viper.GetString("port")
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	fmt.Println("PORT:", port)
	err = srv.Run(port, handlers.InitRoutes())

	if err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigFile("./configs/config.json")
	return viper.ReadInConfig()
}
