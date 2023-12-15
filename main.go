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

	homePage := func(w http.ResponseWriter, r *http.Request) {
		page := template.Must(template.ParseFiles("index.html"))

		pets := map[string][]Pet{
			"Pets": {
				{PetType: "Cat", Color: "orange"},
				{PetType: "Dog", Color: "black"},
				{PetType: "Cat", Color: "orange & white"},
			},
		}

		page.Execute(w, pets)
	}

	addPet := func(w http.ResponseWriter, r *http.Request) {
		log.Print("htmx received")
		log.Print(r.Header.Get("HX-Request"))
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-pet/", addPet)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
