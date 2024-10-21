package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	web "asciiArtWeb/asciiArtFs"
)

type Data struct {
	Text     string
	Banner   string
	AsciiArt template.HTML
}

func main() {
	// Register the handler for the root URL
	http.HandleFunc("/", AppHandler)

	// Start the web server
	log.Println("Starting server on http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil) // use default multiplexer
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func AsciiArtMaker(text string, banner string) (string, []error) {
	errs := []error{}
	if banner == "all" {
		AsciiArt1, err := web.AsciiArtFs(text, "standard")
		errs = append(errs, err)
		AsciiArt2, err := web.AsciiArtFs(text, "shadow")
		errs = append(errs, err)
		AsciiArt3, err := web.AsciiArtFs(text, "thinkertoy")
		errs = append(errs, err)
		return AsciiArt1 + AsciiArt2 + AsciiArt3, errs
	}
	AsciiArt, err := web.AsciiArtFs(text, banner)
	return AsciiArt, []error{err}
}

func AppHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Get(w, r)
	case "POST":
		Post(w, r)
	default:
		http.Error(w, "405 - Method Not Allowed", 405)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Not Found", 404)
		return
	}
	// http.ServeFile(w, r, "template.html")
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "404 - Not Found", 404)
		return
	}
	data := Data{}
	tmpl.Execute(w, data)
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404- Not Found", 404)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "500 - Internal Server Error", 500)
		log.Fatalln(err)
	}

	data := Data{}
	data.Text = r.Form["text"][0]
	data.Banner = r.Form["banner"][0]
	if len(data.Text) == 0 || len(data.Banner) == 0 {
		http.Error(w, "400 - Bad Request", http.StatusBadRequest)
		return
	}

	asciiArt, errs := AsciiArtMaker(data.Text, data.Banner)
	tmpl, err := template.ParseFiles("template.html")
	errs = append(errs, err)

	// Handing template err and AsciiConverter errs
	for i := range errs {
		if errs[i] != nil {
			notFound := fmt.Errorf("NotFound")
			if errs[i] == notFound {
				http.Error(w, "404 - Not Found", 404)
			} else {
				http.Error(w, "500 - Internal Server Error", 500)
			}
			return
		}
	}

	data.AsciiArt = template.HTML(asciiArt)
	tmpl.Execute(w, data)
}
