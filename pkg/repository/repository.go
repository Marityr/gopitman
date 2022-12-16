package repository

import (
	"github.com/Marityr/gopitman"
	"github.com/jmoiron/sqlx"
)

type Authorazation interface {
	CreateUser(user gopitman.User) (int, error)
	GetUser(username, password string) (gopitman.User, error)
}

type Essay interface {
	Create(input gopitman.Essay) error
}

type Repository struct {
	Authorazation
	Essay
}

func NewReposiroty(db *sqlx.DB) *Repository {
	return &Repository{
		Authorazation: NewAuthPostgres(db),
		Essay:         NewEssayPostgres(db),
	}
}
