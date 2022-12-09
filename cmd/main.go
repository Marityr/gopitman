package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/docs"
	"github.com/Marityr/gopitman/pkg/handler"
	"github.com/Marityr/gopitman/pkg/repository"
	"github.com/Marityr/gopitman/pkg/service"
	"github.com/spf13/viper"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in   header
// @name Authorization

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(viper.GetViper())
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewReposiroty(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(gopitman.Server)
	runHttp(srv, handler)

}

func runHttp(srv *gopitman.Server, handler *handler.Handler) {
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("BonusApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("BonusApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Gopitman API"
	docs.SwaggerInfo.Description = "Service swagger api examples"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
