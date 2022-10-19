package repository

import (
	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/pkg/logging"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db      *gorm.DB
	logging *logging.Logger
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user gopitman.User) (int, error) {
	q := DB
	err := q.Create(&user)
	if err.RowsAffected == 0 {
		r.logging.Info(err)
	}
	return user.Id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (gopitman.User, error) {
	var user gopitman.User
	q := DB

	//TODO проверить обработку ошибок от GORM
	if err := q.Where(`username = ? and password = ?`, username, password).First(&user).Error; err != nil {
		r.logging.Info(err)
		return user, err
	}

	return user, nil
}
