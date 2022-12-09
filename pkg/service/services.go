package service

import (
	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/pkg/repository"
	"github.com/gofrs/uuid"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user gopitman.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Customer interface {
	Create(firstName, lastName, birthday, referrerCode, phone, email string) (string, error)
	GetAll(page, limit int) ([]gopitman.Customer, error)
	GetById(id uuid.UUID) (gopitman.Customer, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, input gopitman.UpdateCustomer) error
}

type Service struct {
	Authorization
	Customer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorazation),
		Customer:      NewCustomerService(repos.Customer),
	}
}
