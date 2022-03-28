package main

import (
	"bufio"
	"bytes"
	"html/template"
	"os"
)

type product struct {
	Img         string
	Name        string
	Price       string
	Stars       float64
	Reviews     int
	Description string
}

func subtr(a, b float64) float64 {
	return a - b
}

func list(e ...float64) []float64 {
	return e
}

func main() {

	data := []product{
		{"images/1.png", "strawberries", "$2.00", 4.0, 251, "Description for strawberries"},
		{"images/2.png", "onions", "$2.80", 5.0, 123, "Description for onions"},
		{"images/3.png", "tomatoes", "$3.10", 4.5, 235, "Description for tomatoes"},
		{"images/4.png", "courgette", "$1.20", 4.0, 251, "Description for courgette"},
		{"images/5.png", "broccoli", "$3.80", 3.5, 1230, "Description for broccoli"},
		{"images/6.png", "potatoes", "$3.00", 2.5, 235, "Description for potatoes"},
	}

	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "./templates/"+tmpl)
	}

	templates := template.Must(template.New("").Funcs(template.FuncMap{"subtr": subtr, "list": list}).ParseFiles(allPaths...))

	var processed bytes.Buffer
	templates.ExecuteTemplate(&processed, "page", data)

	outputPath := "./static/index.html"
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()

}
