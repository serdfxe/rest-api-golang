package main

import (
	"log"

	"github.com/serdfxe/rest-api-golang"
	"github.com/serdfxe/rest-api-golang/pkg/handler"
	"github.com/serdfxe/rest-api-golang/pkg/repository"
	"github.com/serdfxe/rest-api-golang/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("Error while init configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(rest.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
