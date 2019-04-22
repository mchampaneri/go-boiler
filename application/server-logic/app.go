package main

import (
	"log"
)

func main() {

	// loading application configurations
	if err := loadAppConfig(); err != nil {
		log.Fatalln("Failed to load config ", err.Error())
	}
	log.Println("Success: Config loaded")

	// registreing web routes for webapp / port 80
	registerWebRoutes()
	log.Println("Success: Routes loaded")
}
