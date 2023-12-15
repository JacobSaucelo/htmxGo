package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
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
		time.Sleep(2 * time.Second)
		petType := r.PostFormValue("pet-type")
		petColor := r.PostFormValue("pet-color")
		page := template.Must(template.ParseFiles("index.html"))
		page.ExecuteTemplate(w, "pet-list-element", Pet{
			PetType: petType,
			Color:   petColor,
		})
		// htmlVal := fmt.Sprintf("<p>Pet: %s - Color: %s</p>", petType, petColor)
		// page, _ := template.New("p").Parse(htmlVal)
		// page.Execute(w, nil)

	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-pet/", addPet)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
