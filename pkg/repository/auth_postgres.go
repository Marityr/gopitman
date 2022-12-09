package repository

import (
	"fmt"
	"log"

	"github.com/Marityr/gopitman"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user gopitman.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (gopitman.User, error) {
	var user gopitman.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		log.Println(err)
	}

	return user, err
}
