package main

import (
	"log"

	restapi "github.com/Bernar11296/rest-api"
	"github.com/Bernar11296/rest-api/pkg/handler"
	"github.com/Bernar11296/rest-api/pkg/repository"
	"github.com/Bernar11296/rest-api/pkg/service"
)

func main() {
	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)
	srv := new(restapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
