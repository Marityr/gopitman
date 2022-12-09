package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(v *viper.Viper) (*sqlx.DB, error) {
	cfg := &Config{
		Host:     v.GetString("db.host"),
		Port:     v.GetString("db.port"),
		Username: v.GetString("db.username"),
		DBName:   v.GetString("db.dbname"),
		SSLMode:  v.GetString("db.sslmode"),
		Password: v.GetString("db.password"),
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return db, nil
}
