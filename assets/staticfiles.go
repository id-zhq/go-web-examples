package main

import "net/http"

/**
http://localhost:8080/static/css/style.css
*/

func main() {

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

}
