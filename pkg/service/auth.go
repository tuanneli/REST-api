package service

import "github.com/tuanneli/REST-api.git/pkg/repository"

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}
