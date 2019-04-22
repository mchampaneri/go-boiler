package main

// Config is structure that provides the configuration
//  parameters to the other parts of the app during the
//  run time
//
var Config struct {
	Env         string `json:"Env"`
	AppName     string `json:"AppName"`
	AppURL      string `json:"AppUrl"`
	Port        string `json:"Port"`
	ViewPath    string `json:"ViewPath"`
	StoragePath string `json:"StoragePath"`
	PublicPath  string `json:"PublicPath"`
}

// StaticPages Struct Used to load the
// Bunch of the static pages
var StaticPages struct {
	Pages []*Page `json:"pages"`
}

// Page show the static page ie. html page
// which are located at the @View  path
// and needed to load over @Url url
type Page struct {
	URL  string `json:"url"`
	View string `json:"view"`
}
