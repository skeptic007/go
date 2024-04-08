package main

import (
	"fmt"
	"log"
)

func main() {
	s, err := NewPostgresStore()
	if err != nil {
		log.Panic(err)
	}

	if err := s.Init(); err != nil {
		fmt.Println("Error while creating table", err)
	}

	server := StartAPI(":6000", s)
	server.Run()

	fmt.Println("Server is listening at port :6000")

}
