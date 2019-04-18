package main

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

// Reading configurations
// and loading static routes
func loadAppConfig() {

	c, cerr := os.Open("./config/app.json")    // Reading Configuration file
	s, serr := os.Open("./config/static.json") // Reading Static routes file

	defer func() {
		c.Close()
		s.Close()
	}()

	configDecoder := json.NewDecoder(c)
	configDecoder.Decode(&Config)

	staticRoutesDecoder := json.NewDecoder(s)
	staticRoutesDecoder.Decode(&StaticPages)

	color.Green(" * Configurations Loaded SuccessFully ")
}
