package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *NewApiStarter) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := s.store.GetUsers()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	WriteJSON(w, http.StatusOK, users)

}

func (s *NewApiStarter) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
	}
	user, err := s.store.GetUser(id)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
	}
	WriteJSON(w, http.StatusOK, user)
}

func (s *NewApiStarter) CreateUserByInput(w http.ResponseWriter, r *http.Request) {
	req := new(User)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
	}
	if err := s.store.CreateUser(req); err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}
	WriteJSON(w, http.StatusOK, req)

}

func (s *NewApiStarter) CreateUserOrGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s.GetUserById(w, r)
	}
	if r.Method == "POST" {
		s.CreateUserByInput(w, r)
	}
	WriteJSON(w, http.StatusMethodNotAllowed, r.Method)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}
