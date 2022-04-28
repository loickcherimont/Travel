package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

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
// Fetch a specific destination by its city name
func GetDestinationByCity(cityName string) []models.Destination {

	// Place to store the specific destination
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

	for _, v := range destinations {
		if v.City == cityName {
			destinations = make([]models.Destination, 1)
			destinations[0] = v
			return destinations
		}
	}

	return make([]models.Destination, 0)
}

// Create a HTML template
func NewHTMLTemplate() *template.Template {
	return template.Must(template.ParseFiles("template/index.html"))
}

// Inject data into the generated template
func GetIndexPage(w http.ResponseWriter, _ *http.Request) {
	tmpl.ExecuteTemplate(w, "Default", GetDestinations())
}

// Fetch user query
// Return the corresponding card
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err.Error())
	}

	params := u.Query()
	city := params.Get("q")

	if r.Method == http.MethodGet {

		// Prepare the user entry
		prepareEntry := func(userEntry string) string {

			// Delete extra spaces
			userEntry = strings.Trim(city, " ")

			// Capitalize city name
			t := strings.Split(userEntry, "")
			firstLetterToUppercase := func(s []string) []string {
				return []string{strings.ToUpper(s[0])}
			}
			t = append(firstLetterToUppercase(t), t[1:]...)
			userEntry = strings.Join(t, "")

			return userEntry
		}

		city = prepareEntry(city)

		// If user entry does not exist
		// Show all available destinations
		if city == "" {
			tmpl.ExecuteTemplate(w, "Default", GetDestinations())
			return
		}
		// Show corresponding card
		tmpl.ExecuteTemplate(w, "Specific", GetDestinationByCity(city))
		return
	}
}
