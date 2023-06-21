package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	fmt.Println("Inside main")
	addr := flag.String("addr", ":2020", "Default port for the server")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := Application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.view)
	mux.HandleFunc("/snippet/create", app.create)

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	err := srv.ListenAndServe()

	if err != nil {
		app.errorLog.Fatal(err)
	}
}
