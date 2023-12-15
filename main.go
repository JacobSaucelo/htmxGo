package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, r.Method)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
