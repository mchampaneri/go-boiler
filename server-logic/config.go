package main

import (
	"github.com/allegro/bigcache"
	"github.com/gorilla/sessions"
)

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
	Database    struct {
		DatabaseName string `json:"DatabaseName"`
		Driver       string `json:"Driver"`
		UserName     string `json:"UserName"`
		Password     string `json:"Password"`
		DatabaseHost string `json:"DatabaseHost"`
	}
	Mail struct {
		Service   string `json:"Service"`
		Domain    string `json:"Domain"`
		Key       string `json:"Key"`
		PublicKey string `json:"PublicKey"`
		APIKey    string `json:"APIKey"`
	}
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

// BigCache //
var BigCache *bigcache.BigCache

// UserSession is the session store which stores the 
// values for the authenticated user
var UserSession = sessions.NewCookieStore([]byte("xf7KylXJ7CFSH4mLZG2Wyl86HAB9Rqvn"))
