package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cuongcb/go-authen/internal/dtos"
	"github.com/cuongcb/go-authen/internal/service"
)

// Register is in charge of recording a new user into system
var Register = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// json.NewEncoder(w).Encode(map[string]string{"error": "not supported method"})
		return
	}

	user := dtos.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	newUser, err := service.CreateUser(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newUser)
}
