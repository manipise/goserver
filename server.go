package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(response, "Server is up and running")
	})
	port := ":8111"
	log.Println("server is listining to port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
