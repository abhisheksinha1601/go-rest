package controllers

import (
	"encoding/json"
	"net/http"
	"rest/bl"
	"rest/interfaces"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	token, err := bl.Login(r.Form.Get("username"), r.Form.Get("password"))
	if err != nil {
		json.NewEncoder(w).Encode(interfaces.IAPIResponse{Error: true, Message: "failure", Result: nil})
		return
	}
	json.NewEncoder(w).Encode(interfaces.IAPIResponse{Error: false, Message: "success", Result: token})
}
