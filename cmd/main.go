package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/IDontCareMe/todo-app"
	"github.com/IDontCareMe/todo-app/pkg/handler"
	"github.com/IDontCareMe/todo-app/pkg/repository"
	"github.com/IDontCareMe/todo-app/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

// initConfig reads config data from configs/config.yaml
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
