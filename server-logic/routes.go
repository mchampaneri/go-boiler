package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// dynamicRoutes requires route
// and handling function for response.
func dynamicRoutes(router *mux.Router) {
	// router.HandleFunc("/example-route", handlingFunction())
}

// staticPagesLoader registers page to main router
func staticPagesLoader(page *Page, nextroute *mux.Route) {
	nextroute.Path(page.URL)
	nextroute.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI, " ", w.Header(), page.View, page.URL)
		View(w, r, nil, page.View)
	})
	return
}
