package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	models "github.com/loickcherimont/Travel/models"
)

var tmpl *template.Template = NewHTMLTemplate()

// From json file
func GetDestinations() []models.Destination {

	// Storage for unmarshal json data
	var destinations []models.Destination

	content, err := os.ReadFile("models/config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &destinations)
	_ = err
	if err != nil {
		log.Fatal(err)
	}

	return destinations
}

// Create a HTML template
func NewHTMLTemplate() *template.Template {
	return template.Must(template.ParseFiles("template/index.html"))
}

// Inject data into the generated template
func GetIndexPage(w http.ResponseWriter, _ *http.Request) {
	tmpl.Execute(w, GetDestinations())
}
