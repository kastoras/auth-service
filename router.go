package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleGetFolder(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Id not provided"))
		return
	}

	res, err := s.storage.getFolder(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Folder not found"))
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	payload := new(CreatePostPayload)

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Invalid payload"))
		return
	}

	fmt.Printf("title = %s", payload.Title)
	fmt.Println()

	folder := &Folder{
		Title:     payload.Title,
		CreatedAt: time.Now(),
	}

	err = s.storage.persistFolder(folder)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Post created"))
}
