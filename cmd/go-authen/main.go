package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cuongcb/go-authen/internal/handler"
	"github.com/cuongcb/go-authen/internal/service"
)

func main() {
	fmt.Println("*** Go Authentication Server ***")

	service.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/user", handler.User)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/register", handler.Register)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Listening on --> localhost:8080")
	log.Fatal(s.ListenAndServe())
}
