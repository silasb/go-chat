package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("chat server")

	r := mux.NewRouter()
	r.HandleFunc("/ws/{id:[0-9]+}", serveWs)

	http.Handle("/", r)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}
}
