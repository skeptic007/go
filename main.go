package main

import (
	"fmt"
	"log"
)

func main() {
	//fmt.Println("hell")
	s, err := NewPostgresStore()
	if err != nil {
		log.Panic(err)
	}

	if err := s.Init(); err != nil {
		fmt.Println("Error while creating table", err)
	}

	server := StartAPI(":8000", s)

	fmt.Println("Server is listening at port :8000")

	server.Run()

}
