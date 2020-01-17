package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cuongcb/go-authen/internal/dtos"
	"github.com/cuongcb/go-authen/internal/service"
)

// User ...
var User = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method != http.MethodGet {
		responseError(w, http.StatusMethodNotAllowed)
		return
	}

	users, err := service.GetUserList()
	if err != nil {
		responseError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*dtos.User{"users": users})
}
