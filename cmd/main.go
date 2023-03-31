package main

import (
	"log"

	restapi "github.com/Bernar11296/rest-api"
	"github.com/Bernar11296/rest-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(restapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
