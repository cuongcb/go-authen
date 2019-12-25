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
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusMethodNotAllowed)})
		return
	}

	user := dtos.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	savedUser, err := service.GetUserByMail(user.Email)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "account not existss"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	if savedUser.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(http.StatusText(http.StatusUnauthorized))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
}
