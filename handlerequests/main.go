package main

import (
	"fmt"
	"net/http"
)

//Use Sigle http Handler interface Starts

type UseHandler struct {
}

func (u *UseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello to Using Handler interface")
}

func fnUseHanderServer() {
	server := http.Server{
		Addr:    ":2020",
		Handler: &UseHandler{},
	}
	server.ListenAndServe()
}

// Use Single http Handler interface Ends

// Use Single http HandlerFunc Starts
func fnUseHandleFuncServer() {
	http.HandleFunc("/", UseHandleFunc)
	http.ListenAndServe(":2020", nil)
}

func UseHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello to using HandleFunc")
}

//Use Single http HandlerFunc Ends

//Use Multiple Handler interface Starts

type UseHandler1 struct {
}

func (u1 *UseHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Use multiple Handler interface - this is handler 1")
}

type UseHandler2 struct {
}

func (u2 UseHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Use multiple Handler interface - this is handler 2")
}

func fnUseMultipleHandlerServer() {
	http.Handle("/handler1", &UseHandler1{})
	http.Handle("/handler2", &UseHandler2{})
	http.ListenAndServe(":2020", nil)
}

//Use Multiple Handler interface Ends

//Use Chain Handler interface Starts

type UseChainHandler struct {
}

func (u UseChainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Use chaining Handlers - You are inside main handler")
}

func ChainHandler1(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Use chaining Handlers - You are inside Chain Handler 1")
		handler.ServeHTTP(w, r)
	})
}

func ChainHandler2(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Use chaining Handlers - You are inside Chain Handler 2")
		handler.ServeHTTP(w, r)
	})
}

func fnUseChainHandlerServer() {
	handler := UseChainHandler{}
	http.Handle("/chainhandler", ChainHandler2(ChainHandler1(handler)))
	http.ListenAndServe(":2020", nil)
}

//Use Chain Handler interface Ends

//Use Chain HandlerFunc Starts

func ChainHandlerFunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Chain HandlerFunc1")
}

func ChainHandlerFunc2(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside Chain HandlerFunc2")
		h(w, r)
	}
}

func fnUseChainHandlerFunc() {
	http.HandleFunc("/chainhandlerfunc", ChainHandlerFunc2(ChainHandlerFunc1))
	http.ListenAndServe(":2020", nil)
}

//Use Chain HandlerFunc Ends

func main() {
	//fnUseHanderServer()
	//fnUseHandleFuncServer()
	//fnUseMultipleHandlerServer()
	//fnUseChainHandlerServer()
	fnUseChainHandlerFunc()
}
