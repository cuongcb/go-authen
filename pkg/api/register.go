package api

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github.com/cuongcb/go-authen/pkg/service"
	"net/http"
)

// Register is in charge of recording a new user into system
var Register = func(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err)
	}

	if err := service.CreateUser(user.Email, user.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "internal server error")
	}
}
