package service

import (
	"github.com/IDontCareMe/todo-app"
	"github.com/IDontCareMe/todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func newAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	return s.repo.CreateUser(user)
}
