package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cuongcb/go-authen/internal/dtos"
	"github.com/cuongcb/go-authen/internal/service"
	"github.com/cuongcb/go-authen/internal/service/log"
)

// Register is in charge of recording a new user into system
var Register = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	user := dtos.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Log(err)
		responseError(w, http.StatusBadRequest)
		return
	}

	newUser, err := service.CreateUser(user.Email, user.Password)
	if err != nil {
		log.Log(err)
		responseError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
