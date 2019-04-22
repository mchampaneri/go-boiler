package main

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
)

// Reading configurations
// and loading static routes
func loadAppConfig() error {

	c, cerr := os.Open("./config/app.json") // Reading Configuration file
	if cerr != nil {
		return cerr
	}
	s, serr := os.Open("./config/static.json") // Reading Static routes file
	if serr != nil {
		return serr
	}

	defer func() {
		c.Close()
		s.Close()
	}()

	configDecoder := json.NewDecoder(c)
	configDecoder.Decode(&Config)

	staticRoutesDecoder := json.NewDecoder(s)
	staticRoutesDecoder.Decode(&StaticPages)

	color.Green(" * Configurations Loaded SuccessFully ")
	return nil
}
