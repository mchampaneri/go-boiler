package main

import (
	"github.com/allegro/bigcache"
	"github.com/gorilla/sessions"
)

/* Config is structure that provides the configuration
|  parameters to the other parts of the app during the
|  run time
*/
var Config struct {
	Env         string `json:"Env"`
	AppName     string `json:"AppName"`
	AppUrl      string `json:"AppUrl"`
	Port        string `json:"Port"`
	ViewPath    string `json:"ViewPath"`
	StoragePath string `json:"StoragePath"`
	PublicPath  string `json:"PublicPath"`
	Database struct {
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
	}
}
var StaticPages struct {
	Pages []*Page `json:"pages"`
}
type Page struct {
	Url  string `json:"url"`
	View string `json:"view"`
}

// Big Cache //
var BigCache *bigcache.BigCache

// Session Manager //
var UserSession = sessions.NewCookieStore([]byte("xf7KylXJ7CFSH4mLZG2Wyl86HAB9Rqvn"))

