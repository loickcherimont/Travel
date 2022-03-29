package main

import (
	"html/template"
	"log"
	"net/http"
)

type Destination struct {
	Image       string
	AltImage    string
	City        string
	Country     string
	Continent 	string
	Stars       float64
	Description string
}

// Prepare data for HTML template
var data = []Destination{
	{"sanfrancisco.jpg", "Golden Gate Bridge", "San Francisco", "USA", "America", 4, "Lorem ipsum"},
	{"paris.jpg", "Eiffel Tower", "Paris", "France", "Eurasia", 4.5, "Lorem ipsum"},
	{"auckland.jpg", "Auckland downtown", "Auckland", "New Zealand", "Oceania", 4.75, "Lorem ipsum"},
}

// Generate the template
var tmpl *template.Template = template.Must(template.ParseFiles("template/index.html"))

// Handle my template
func homeHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl.Execute(w, data)
}

func main() {

	// Serve static files in directory for assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Manage and Listen and serve
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
