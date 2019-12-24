package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cuongcb/go-authen/pkg/api"
	"github.com/cuongcb/go-authen/pkg/service"
)

func main() {
	fmt.Println("*** Go Authentication Server ***")

	service.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Home)
	mux.HandleFunc("/register", api.Register)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Listening on: 127.0.0.1:8080")
	log.Fatal(s.ListenAndServe())
}
