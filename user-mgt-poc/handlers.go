package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var storage = NewStorage()

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello")
}

// User handler that handles POST/Users
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := storage.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}