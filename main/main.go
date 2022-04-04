package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/loickcherimont/Travel/handler"
)

func main() {
	// Serve CSS, Javascript and images (static files) (1)
	fs := http.FileServer(http.Dir("assets"))

	// Handle index page with styles
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs)) // (2)
	mux.HandleFunc("/travel", handler.GetIndexPage)

	// Serve on 8080 PORT
	fmt.Println("Listening and serving on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
