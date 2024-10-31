package main

import (
	"html/template"
	"log"
	"net/http"
	"asciiArtWeb/httpHandlers"
)

type Data struct {
	Text     string
	Banner   string
	AsciiArt template.HTML
}

func main() {
	// Register the handler for the root URL
	http.HandleFunc("/", httpHandlers.AppHandler)

	// Start the web server
	log.Println("Starting server on http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil) // use default multiplexer
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
