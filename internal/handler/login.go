package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cuongcb/go-authen/internal/dtos"
	"github.com/cuongcb/go-authen/internal/service"
)

// Login ...
var Login = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method != http.MethodPost {
		responseError(w, http.StatusMethodNotAllowed)
		return
	}

	user := dtos.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responseError(w, http.StatusBadRequest)
		return
	}

	savedUser, err := service.VerifyUser(user.Email, user.Password)
	if err == sql.ErrNoRows {
		responseError(w, http.StatusInternalServerError)
		return
	}

	if err != nil {
		// TODO(cuongcb): check server errors later
		// responseError(w, http.StatusInternalServerError)
		responseError(w, http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedUser)
}
