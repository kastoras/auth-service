package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addres string
	client *Client
}

func NewAPIServer(addres string) *APIServer {
	httpClient := &http.Client{}
	return &APIServer{
		addres: addres,
		client: &Client{
			httpClient: httpClient,
		},
	}
}

func (s *APIServer) Run() {

	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet).Methods("GET")
	router.HandleFunc("/login", s.handleLogin).Methods("POST")

	http.ListenAndServe(s.addres, router)
	fmt.Println("Server started...")
}
