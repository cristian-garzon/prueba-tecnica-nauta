package main

import (
	"log"

	"prueba-tecnica-nauta/app/infrastructure/builder"
	"prueba-tecnica-nauta/app/infrastructure/config"
	"prueba-tecnica-nauta/app/infrastructure/server"
)

func main() {
	log.Println("Starting reader service...")

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	repositoriesBuilder, err := builder.NewRepositoriesBuilder(config)
	if err != nil {
		log.Fatal(err)
	}

	actionsBuilder := builder.NewActionsBuilder(repositoriesBuilder)
	server := server.SetupServer(config.Server, actionsBuilder)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
