package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, nil)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}

func main() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8080", nil)

}
