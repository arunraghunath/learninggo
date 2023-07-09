package main

import (
	"html/template"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	files := []string{"./templates/layout.html", "./templates/public.navbar.html", "./templates/signup.html"}
	ts := template.Must(template.ParseFiles(files...))

	ts.ExecuteTemplate(w, "layout", nil)
}
