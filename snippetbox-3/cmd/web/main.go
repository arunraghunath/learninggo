package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Inside main")

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", view)
	mux.HandleFunc("/snippet/create", create)

	err := http.ListenAndServe(":2020", mux)

	if err != nil {
		log.Fatal(err)
	}

}
