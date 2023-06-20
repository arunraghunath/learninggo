package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/pages/base.html",
		"ui/html/pages/home.html",
		"ui/html/partials/nav.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errLog.Print(err)
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errLog.Print(err)
	}
}

func (app *Application) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is not allowed", 405)
	}
	w.Write([]byte("Create a new Snippet Box"))
}

func (app *Application) view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	app.infoLog.Printf("Display the content of a specific snippet box with ID: %d", id)
}
