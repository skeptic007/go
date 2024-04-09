package main

import (
	"net/http"
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
	//r := mux.NewRouter()
	//r.HandleFunc("/users", s.GetAllUsers)
	//r.HandleFunc("/user", s.CreateUserOrGetUser).Methods("GET", "POST")
	//r.HandleFunc("/user/:id", s.GetUserById).Methods("GET")
	http.HandleFunc("/bar", s.GetAllUsers)

	http.ListenAndServe(s.Addr, nil)
}
