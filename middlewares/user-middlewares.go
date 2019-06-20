package middlewares

import (
	"net/http"
	"rest/controllers"

	"github.com/gorilla/mux"
)

func UserMiddleware(next http.Handler) http.Handler {
	app := mux.NewRouter()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.HandleFunc("users/login", controllers.LoginController)
		// next.ServeHTTP(w, r)
	})
}
