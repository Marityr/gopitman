package repository

import (
	"fmt"

	"github.com/Marityr/gopitman/pkg/logging"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(v *viper.Viper) (*gorm.DB, error) {
	logger := logging.GetLooger()

	cfg := &Config{
		Host:     v.GetString("db.host"),
		Port:     v.GetString("db.port"),
		Username: v.GetString("db.username"),
		DBName:   v.GetString("db.dbname"),
		SSLMode:  v.GetString("db.sslmode"),
		Password: v.GetString("db.password"),
	}

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		LogLevel: logger.Info, // Log level
	// 		Colorful: true,        // Disable color
	// 	},
	// )

	bURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.SSLMode, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(postgres.Open(bURL), &gorm.Config{
		// отключение зависимости связей на уровне базы
		DisableForeignKeyConstraintWhenMigrating: false,
		// кеширование запроса
		PrepareStmt: true,
		// Logger:      newLogger,
	})

	if err != nil {
		logger.Fatal(err)
	}

	DB = db

	return db, nil
}
