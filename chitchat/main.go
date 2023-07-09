package main

import (
	"fmt"
	"net/http"
)

func main() {
	p("Inside func main for Chitchat", version(), "hi")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/", index)

	http.ListenAndServe(":2020", mux)

}

func p(a ...interface{}) {
	fmt.Println(a...)
}

func version() string {
	return "0.1"
}
