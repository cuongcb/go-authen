package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cuongcb/go-authen/internal/dtos"
	"github.com/cuongcb/go-authen/internal/service"
)

// Home ...
var Home = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "not supported method"})

		return
	}

	users, err := service.GetUserList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*dtos.User{"users": users})
}
