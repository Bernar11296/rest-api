package main

import (
	"log"

	restapi "github.com/Bernar11296/rest-api"
	"github.com/Bernar11296/rest-api/pkg/handler"
	"github.com/Bernar11296/rest-api/pkg/repository"
	"github.com/Bernar11296/rest-api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err:=initConfig(); err!=nil{
		log.Fatal("error initializing configs: %s", err.Error())
	}
	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)
	srv := new(restapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}