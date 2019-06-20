package main

import (
	"fmt"
	"log"
	"net/http"
	"rest/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server...")

	app := mux.NewRouter().StrictSlash(true)
	app.Use(middlewares.AppMiddleware)
	app.PathPrefix("/users/").Subrouter().Use(middlewares.UserMiddleware)

	http.Handle("/", app)
	log.Fatal(http.ListenAndServe(":4000", app))
}
