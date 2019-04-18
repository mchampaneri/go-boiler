package main

import "github.com/fatih/color"

func main() {

	// //Loading the Config
	loadAppConfig()
	color.Yellow(" * %s : Powred by go-boiler  ", Config.AppName)

	// if dbErr := initdb(); dbErr != nil {
	// 	color.Red("* Database: failed %s ", dbErr.Error())
	// 	return
	// }

	// // Enforcing the goValidator over the models (Structs)
	// govalidator.SetFieldsRequiredByDefault(false)
	// DefaultLogger = Log{}

	// // Using BigCache as default cache
	// BigCache, _ = initCache()

	// // Loading the Routers for the web and api on their
	// // mentioned ports
	// //
	// RegisterWebRoutes()
}
