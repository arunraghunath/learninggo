package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {

	fmt.Println("Inside main")

	addr := flag.String("addr", ":2020", "Default port for starting the server")
	flag.Parse()
	mux := http.NewServeMux()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ldate)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &Application{infoLog: infoLog, errLog: errLog}

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.create)
	mux.HandleFunc("/snippet/view", app.view)

	infoLog.Print("Before starting server at", *addr)
	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
