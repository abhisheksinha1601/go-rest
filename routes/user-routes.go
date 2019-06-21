package routes

import (
	"rest/controllers"

	"github.com/gorilla/mux"
)

// func UserMiddleware(next http.Handler) http.Handler {
// 	app := mux.NewRouter()
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("User middleware")
// 		app.HandleFunc("/users/login", controllers.LoginController)
// 		// next.ServeHTTP(w, r)
// 	})
// }

func UserRoutes(app *mux.Router) {
	app.HandleFunc("/login", controllers.LoginController)
}
