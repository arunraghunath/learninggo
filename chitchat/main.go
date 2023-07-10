package main

import (
	"fmt"
	"net/http"
)

func main() {

	p("Inside func main for Chitchat", version(), "hi1")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	p("Inside func main for Chitchat", version(), "hi2")

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/authenticate", authenticate)
	p("Inside func main for Chitchat", version(), "hi3")

	err := http.ListenAndServe(":2021", mux)
	if err != nil {
		fmt.Println(err)
	}
	p("Inside func main for Chitchat", version(), "hi3")

}

func p(a ...interface{}) {
	fmt.Println(a...)
}

func version() string {
	return "0.1"
}
