package main

import (
	"fmt"
	"net/http"
)

/**
展示index.html
http://localhost/static/

展示文字“Welcome to my website!”
http://localhost/
*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
