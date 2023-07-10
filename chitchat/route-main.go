package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/arunraghunath/learninggo/chitchat/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	thread, err := models.Threads()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = session(w, r)

	files := []string{"./templates/layout.html", "./templates/public.navbar.html", "./templates/index.html"}
	if err == nil {
		files = []string{"./templates/layout.html", "./templates/private.navbar.html", "./templates/index.html"}
	}
	ts := template.Must(template.ParseFiles(files...))
	ts.ExecuteTemplate(w, "layout", thread)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		fmt.Println("Cookie does not exist")
	} else {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid Session")

		}
	}
	return

}
