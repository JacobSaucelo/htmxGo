package main

import (
	"log"
	"net/http"
	"text/template"
)

type Pet struct {
	PetType string
	Color   string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := template.Must(template.ParseFiles("index.html"))

		pets := map[string][]Pet{
			"Pets": {
				{PetType: "Cat", Color: "orange"},
				{PetType: "Dog", Color: "black"},
				{PetType: "Cat", Color: "orange & white"},
			},
		}

		page.Execute(w, pets)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
