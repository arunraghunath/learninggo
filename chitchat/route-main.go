package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"./templates/layout.html", "./templates/public.navbar.html", "./templates/index.html"}
	ts := template.Must(template.ParseFiles(files...))
	ts.ExecuteTemplate(w, "layout", nil)
}
