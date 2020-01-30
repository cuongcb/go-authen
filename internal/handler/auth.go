package handler

import (
	"net/http"

	"github.com/cuongcb/go-authen/internal/service"
)

// AuthMw ...
var AuthMw = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.FormValue("session")
		if _, err := service.Session(session); err != nil {
			return
		}

		next.ServeHTTP(w, r)
	})
}
