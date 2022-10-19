package repository

import (
	"github.com/Marityr/gopitman"
	"gorm.io/gorm"
)

type Authorazation interface {
	CreateUser(user gopitman.User) (int, error)
	GetUser(username, password string) (gopitman.User, error)
}
type Customer interface{}
type Transaction interface{}

type Repository struct {
	Authorazation
	Customer
	Transaction
}

func NewReposiroty(db *gorm.DB) *Repository {
	return &Repository{
		Authorazation: NewAuthPostgres(db),
	}
}
