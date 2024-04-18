package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type NewApiStarter struct {
	Addr  string
	store *PostgresStore
}

func StartAPI(addr string, store *PostgresStore) *NewApiStarter {
	return &NewApiStarter{
		Addr:  addr,
		store: store,
	}
}

func (s *NewApiStarter) Run() {
	r := mux.NewRouter()
	//fmt.Println("reached")

	//r.HandleFunc("/users", s.GetAllUsers)
	r.HandleFunc("/user", s.CreateUserOrGetUser).Methods("GET", "POST")
	//r.HandleFunc("/user/:id", s.GetUserById).Methods("GET")
	r.HandleFunc("/bar", s.GetAllUsers).Methods("GET")
	fmt.Println("reached")

	http.ListenAndServe(s.Addr, r)
}
