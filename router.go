package main

import (
	"encoding/json"
	"net/http"
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

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {

	payload := new(LoginPayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid payload"))
	}

	kpayload := &KloginPayload{
		clientID:     "authentication-api-auth",
		username:     payload.Username,
		password:     payload.Password,
		grandType:    "password",
		clientSecret: "1JUmMOVpDVGssWNvJvOEEfgXa6qVvMOU",
	}

	kloginresp, err := s.client.login(kpayload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res := &LoginResp{
		AccessToken: kloginresp.AccessToken,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
