package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book %s on page %s\n", title, page)
	})

	r.HandleFunc("/books/{title}", CreateBook()).Methods("POST")
	r.HandleFunc("/books/{title}", CreateBook()).Methods("GET")
	r.HandleFunc("/books/{title}", CreateBook()).Methods("PUT")
	r.HandleFunc("/books/{title}", CreateBook()).Methods("DELETE")

	r.HandleFunc("/books/{title}", CreateBook()).Host("www.mybookstore.com")

	r.HandleFunc("/secure", CreateBook()).Schemes("https")
	r.HandleFunc("/insecure", CreateBook()).Schemes("http")

	bookroute := r.PathPrefix("/books").Subrouter()
	bookroute.HandleFunc("/", CreateBook())
	bookroute.HandleFunc("/{title}", CreateBook())

	http.ListenAndServe(":80", r)
}

func CreateBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
