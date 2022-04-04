package handler

import (
	"html/template"
	"net/http"

	models "github.com/loickcherimont/Travel/models"
)

var tmpl *template.Template = NewHTMLTemplate()

// TODO: Use JSON file to stock data
// instead of the following function
func GetInfoOnDestinations() []models.Destination {
	destinations := []models.Destination{
		{"sanfrancisco.jpg", "Golden Gate Bridge", "San Francisco", "USA", "America", 4, "Lorem ipsum"},
		{"paris.jpg", "Eiffel Tower", "Paris", "France", "Eurasia", 4.5, "Lorem ipsum"},
		{"auckland.jpg", "Auckland downtown", "Auckland", "New Zealand", "Oceania", 4.5, "Lorem ipsum"},
	}

	return destinations
}

// Create a HTML template
func NewHTMLTemplate() *template.Template {
	return template.Must(template.ParseFiles("template/index.html"))
}

// Inject data into the generated template
func GetIndexPage(w http.ResponseWriter, _ *http.Request) {
	tmpl.Execute(w, GetInfoOnDestinations())
}
