package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/arunraghunath/learninggo/chitchat/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	files := []string{"./templates/layout.html", "./templates/public.navbar.html", "./templates/signup.html"}
	ts := template.Must(template.ParseFiles(files...))

	ts.ExecuteTemplate(w, "layout", nil)
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	user := &models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	err = user.Create()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Could not create user")
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	files := []string{"./templates/layout.html", "./templates/public.navbar.html", "./templates/login.html"}
	ts := template.Must(template.ParseFiles(files...))
	ts.ExecuteTemplate(w, "layout", nil)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("User does not exist"))
	}
	if user.Password != models.Encrypt(r.PostFormValue("password")) {
		w.Write([]byte("Incorrect credentials"))
		http.Redirect(w, r, "/login", 302)
	} else {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	}
}
