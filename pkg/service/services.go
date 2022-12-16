package service

import (
	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user gopitman.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Essay interface {
	Create(input gopitman.Essay) error
}

type Service struct {
	Authorization

	Essay
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorazation),
		Essay:         NewEssayService(repos.Essay),
	}
}
