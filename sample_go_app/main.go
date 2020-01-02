package main

import (
	"log"
	"net/http"
)

type IndexHandler struct{}

type AppHandler struct{}

type DefaultHandler struct{}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	<h1>Hello World!!</h1>
	`))
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	<h1>Your Sample App!!</h1>
	`))
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URI: %v", r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	http.Handle("/index.html", &IndexHandler{})
	http.Handle("/app/", &AppHandler{})
	http.Handle("/", &DefaultHandler{})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
