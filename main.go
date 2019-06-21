package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest/interfaces"
	"rest/modules/common"
	"rest/modules/mongo"
	"rest/routes"

	"github.com/gorilla/mux"
)

var openRoutes = []string{
	"/users/login",
	"/users/forgotPassword",
	"/users/resetPassword",
	"/users/signup",
	"/users/verify",
	"/users/exists",
}

func appMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to: " + r.URL.Path)
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()

		for _, route := range openRoutes {
			if route == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}
		common.ParseTokenFromHeader(r)
		if r.Header.Get("user") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(interfaces.IAPIResponse{Error: true, Message: "unauthorized", Result: false})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Starting server...")

	app := mux.NewRouter().StrictSlash(true)

	app.Use(appMiddleware)

	routes.UserRoutes(app.PathPrefix("/users").Subrouter())
	http.Handle("/", app)

	mongo.GetMongoDb()

	log.Fatal(http.ListenAndServe(":4000", app))
}
