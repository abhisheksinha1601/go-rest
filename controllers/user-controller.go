package controllers

import (
	"encoding/json"
	"net/http"
	"rest/interfaces"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(interfaces.IAPIResponse{Error: false, Message: "success", Result: true})
}
