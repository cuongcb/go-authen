package main

import (
	"fmt"
	"github.com/cuongcb/go-authen/pkg/api"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Authentication Starting...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Home)
	mux.HandleFunc("/register", api.Register)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
