package main

import (
	"fmt"
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
		petType := r.PostFormValue("pet-type")
		petColor := r.PostFormValue("pet-color")
		htmlVal := fmt.Sprintf("<p>Pet: %s - Color: %s</p>", petType, petColor)
		page, _ := template.New("p").Parse(htmlVal)
		page.Execute(w, nil)

	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-pet/", addPet)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
