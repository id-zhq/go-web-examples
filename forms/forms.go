package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms/forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		fmt.Printf("recived message : email = %s, subject = %s, message = %s\n", details.Email, details.Subject, details.Message)
		// do someting with details
		_ = details

		tmpl.Execute(w, struct {
			Success bool
		}{true})
	})

	http.ListenAndServe(":8080", nil)
}
