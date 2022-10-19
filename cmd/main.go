package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/pkg/handler"
	"github.com/Marityr/gopitman/pkg/logging"
	"github.com/Marityr/gopitman/pkg/repository"
	"github.com/Marityr/gopitman/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/swaggo/swag/example/basic/docs"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in   header
// @name Authorization

func main() {
	//TODO настроить вывод стандартных логов в поток сервера
	//logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logging.GetLooger()

	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewPostgresDB(viper.GetViper())
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewReposiroty(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(gopitman.Server)
	runHttp(srv, handler, logger)

}

func runHttp(srv *gopitman.Server, handler *handler.Handler, logger logging.Logger) {
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Print("BonusApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Print("BonusApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Bonus API"
	docs.SwaggerInfo.Description = "Bonus swagger api examples"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
