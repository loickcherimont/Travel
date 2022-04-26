package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	models "github.com/loickcherimont/Travel/models"
)

var tmpl *template.Template = NewHTMLTemplate()

// From json file
// Fetch all destinations
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

// From json file
// Fetch a specific destination by its city
func GetDestinationByCity(cityName string) []models.Destination {

	// Place to store the specific destination
	var destination []models.Destination

	content, err := os.ReadFile("models/config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &destination)
	_ = err
	if err != nil {
		log.Fatal(err)
	}

	// Debug
	// for k, v := range destination {
	// 	if
	// }

	return destination
}

// Create a HTML template
func NewHTMLTemplate() *template.Template {
	return template.Must(template.ParseFiles("template/index.html"))
}

// Inject data into the generated template
func GetIndexPage(w http.ResponseWriter, _ *http.Request) {
	tmpl.ExecuteTemplate(w, "Default", GetDestinations())
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err.Error())
	}

	params := u.Query()
	city := params.Get("q")

	if r.Method == http.MethodGet {
		// user entry is empty
		if city == "" {
			// Show all available destinations
			tmpl.ExecuteTemplate(w, "Default", GetDestinations())
		}
		tmpl.ExecuteTemplate(w, "Specific", GetDestinationByCity(city))
	}
}
