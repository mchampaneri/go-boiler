package main

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

/**
 *  Registering the routes for the web
 */
func registerWebRoutes() {

	mainrouter := mux.NewRouter()

	// Public Folder Handler
	mainrouter.PathPrefix("/public/").Handler(http.StripPrefix("/public/",
		http.FileServer(http.Dir("./public"))))
	// Storage Folder Handler
	mainrouter.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/",
		http.FileServer(http.Dir("./storage"))))
	// Favicon Handler
	mainrouter.HandleFunc("/favicon.ico", faviconHandler)
	// Loading The Static Routes
	for _, page := range StaticPages.Pages {
		currRoute := mainrouter.NewRoute()
		staticPagesLoader(page, currRoute)
	}
	// Loading The Dynamic Routes
	dynamicRoutes(mainrouter)
	// Registring 404/Not Found Handler
	mainrouter.NotFoundHandler = http.HandlerFunc(notFoundHandle)
	http.Handle("/", mainrouter)

	log.Println(Config.Port)

	http.ListenAndServe(Config.Port,
		handlers.CompressHandler(csrf.Protect([]byte("El0a6L8uqv"),
			csrf.Secure(true))(mainrouter)))

}

func notFoundHandle(w http.ResponseWriter, r *http.Request) {
	View(w, r, nil, "404.html")
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}
