package repository

import (
	"github.com/Marityr/gopitman"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type Authorazation interface {
	CreateUser(user gopitman.User) (int, error)
	GetUser(username, password string) (gopitman.User, error)
}

type Customer interface {
	Create(firstName, lastName, birthday, referrerCode, phone, email string) (string, error)
	GetAll(page, limit int) ([]gopitman.Customer, error)
	GetById(id uuid.UUID) (gopitman.Customer, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, input gopitman.UpdateCustomer) error
}

type Repository struct {
	Authorazation
	Customer
}

func NewReposiroty(db *sqlx.DB) *Repository {
	return &Repository{
		Authorazation: NewAuthPostgres(db),
		Customer:      NewCustomerPostgres(db),
	}
}
