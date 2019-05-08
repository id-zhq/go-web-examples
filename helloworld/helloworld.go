package main

import (
	"fmt"
	"net/http"
)

func main() {
	//registering a request handler
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
