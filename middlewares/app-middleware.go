package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest/interfaces"
)

var openRoutes = []string{
	"/users/login",
	"/users/forgotPassword",
	"/users/resetPassword",
	"/users/signup",
	"/users/verify",
	"/users/exists",
}

func AppMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to: " + r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		for _, route := range openRoutes {
			if route == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}
		fmt.Println(r.Header.Get("user"))
		// parseTokenFromHeader(req)
		if r.Header.Get("user") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(interfaces.IAPIResponse{Error: true, Message: "unauthorized", Result: false})
			return
		}
		next.ServeHTTP(w, r)
	})
}
