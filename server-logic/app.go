package main

import (
	"log"
)

func main() {

	if err := loadAppConfig(); err != nil {
		log.Fatalln("Failed to load config ", err.Error())
	}
	log.Println("Success: Config loaded")

	registerWebRoutes()
	log.Println("Success: Routes loaded")
}
