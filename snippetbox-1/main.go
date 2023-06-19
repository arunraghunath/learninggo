package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to snippetbox"))
}

func view(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("id"))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display the content for a specific snippet with ID: %d", id)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	fmt.Println("Inside main")

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view?id=2", view)

	mux.HandleFunc("/snippet/create", create)

	log.Println("Before starting a server at :2020")

	log.Fatal(http.ListenAndServe(":2020", mux))

}
