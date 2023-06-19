package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", view)
	mux.HandleFunc("/snippet/create", create)

	log.Print("Starting server on port 2020")
	err := http.ListenAndServe(":2020", mux)
	log.Fatal(err)
}
