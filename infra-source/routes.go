package main

import "github.com/gorilla/mux"

func dynamicRoutes(router *mux.Router){

	router.HandleFunc("/register",registerUser)
	router.HandleFunc("/users",allUsers)
	router.HandleFunc("/login",loginUser)

}
