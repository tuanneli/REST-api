package service

import (
	todo "github.com/tuanneli/REST-api.git"
	"github.com/tuanneli/REST-api.git/pkg/repository"
)

type Authorization interface {
	createUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
