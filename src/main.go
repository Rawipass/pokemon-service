package main

import (
	"log"

	//	"github.com/Rawipass/pokemon-service/config"
	"github.com/Rawipass/pokemon-service/routes"
)

func main() {
	// Init Config
	//	config.InitConfig()

	// Init Database
	//	config.ConnectDatabase()

	// Setup Route
	r := routes.SetupRouter()
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Could not run server: %v\n", err)
	}

}
