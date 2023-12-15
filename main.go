package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := template.Must(template.ParseFiles("index.html"))
		page.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
