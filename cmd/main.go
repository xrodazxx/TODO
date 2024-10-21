package main

import (
	"log"

	"github.com/xrodazxx/TODO/pkg/handler"
	"github.com/xrodazxx/TODO/tod"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(tod.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server:%s", err.Error())
	}

}
