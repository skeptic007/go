package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ebpearls/gorest/handler"
)

func main() {
	// http.Get("/", func(w http.ResponseWriter, err error) {
	// 	//fmt.Println("Server is running")

	// })

	http.HandleFunc("/bar", handler.HandleRequest)
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
