package httpHandlers

import (
	"net/http"
	"html/template"
)

type Data struct {
	Text     string
	Banner   string
	AsciiArt template.HTML
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
	
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "404 - Not Found (Could not find template)", 404)
		return
	}

	data := Data{}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 - Internal Server Error", 500)
		return
	}
}