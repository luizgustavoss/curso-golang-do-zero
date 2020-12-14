package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	usuario := Usuario{Nome: "Luiz"}
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/ola", func(w http.ResponseWriter, r *http.Request){
		templates.ExecuteTemplate(w, "home.html", usuario)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}

type Usuario struct {
	Nome string
	Idade uint8
}