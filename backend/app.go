package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/fatih/color"
)

func main() {


	//Loading the Config
	loadAppConfig()
	color.Yellow(" * %s : A Puberstreet Inc. Product ", Config.AppName)

	initdb()
	color.Green(" * %s Database on %s initiated ",Config.Database.DatabaseName, Config.Database.Driver)
	//
	// Enforcing the goValidator over the models (Structs)
	//
	govalidator.SetFieldsRequiredByDefault(false)
	DefaultLogger = Log{}
	// Wiper uses the BigCache as its built in cache service
	// provider . Initializing the  Cache Singleton
	//
	BigCache, _ = initCache()

	// Loading the Routers for the web and api on their
	// mentioned ports
	//
	RegisterWebRoutes()
}
