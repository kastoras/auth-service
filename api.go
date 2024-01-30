package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addres  string
	storage *Storage
}

func NewAPIServer(addres string, db *gorm.DB) *APIServer {
	return &APIServer{
		addres: addres,
		storage: &Storage{
			db: db,
		},
	}
}

func (s *APIServer) Run() {
	s.migrate()

	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet).Methods("GET")
	router.HandleFunc("/folder/{id}", s.handleGetFolder).Methods("GET")
	router.HandleFunc("/folder", s.handleCreatePost).Methods("POST")

	http.ListenAndServe(s.addres, router)
	fmt.Println("Server started...")
}

func (s *APIServer) migrate() {
	s.storage.db.AutoMigrate(&Folder{})
}
