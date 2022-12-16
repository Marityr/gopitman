package repository

import (
	"fmt"
	"log"

	"github.com/Marityr/gopitman"
	"github.com/jmoiron/sqlx"
)

type EssayPostgres struct {
	db *sqlx.DB
}

func NewEssayPostgres(db *sqlx.DB) *EssayPostgres {
	return &EssayPostgres{db: db}
}

func (e *EssayPostgres) Create(input gopitman.Essay) error {
	createEssay := fmt.Sprintf("INSERT INTO %s (title, address, coordinates, descriptions) VALUES ($1, $2, $3, $4)", "essay")
	_, err := e.db.Exec(createEssay, input.Title, input.Address, input.Coordinates, input.Descriptions)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
