package main

import (
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

func AsciiArtMaker(text string, banner string) (string, []any) {
	errs := []any{}
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
	return AsciiArt, []any{err}
}

func AppHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Get(w, r)
	case "POST":
		Post(w, r)
	default:
		http.Error(w, "400 - Bad request (bad method used)", 400)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ascii-art" {
		http.Error(w, "400 - Bad request (bad method used)", 400)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "404 - Not Found", 404)
		return
	}
	// http.ServeFile(w, r, "template.html")
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "404 - Not Found (Could not find template)", 404)
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
	switch {
	case len(data.Text) == 0 || len(data.Banner) == 0:
		http.Error(w, "400 - Bad Request (It's not allowed to keep empty inputs)", 400)
		return
	case len(data.Text) > 500:
		http.Error(w, "400 - Bad Request (length of the text is above 500 char.)", 400)
		return
	}

	asciiArt, errs := AsciiArtMaker(data.Text, data.Banner)
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "404 - Not Found (Could not find template)", 404)
	}

	// Handing template err and AsciiConverter errs	
	for i := range errs {
		if errs[i] != nil {
			if errs[i] == "NotFound" {
				http.Error(w, "404 - Not Found (Could not find banner)", 404)
			} else if errs[i] == "Non-Ascii" {
				http.Error(w, "400 - Bad Request (A none ascii char has been found)", 400)
			} else {
				http.Error(w, "500 - Internal Server Error", 500)
			}
			return
		}
	}

	data.AsciiArt = template.HTML(asciiArt)
	tmpl.Execute(w, data)
}
