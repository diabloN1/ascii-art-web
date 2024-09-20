package main

import (
	"html/template"
	"log"
	"net/http"
	web "asciiArtWeb/asciiArtFs"
)

type Data struct {
	Text string
	Banner string
	AsciiArt string
}

func main() {
    // Register the handler for the root URL
    http.HandleFunc("/", htmlHandler)

    // Start the web server
    log.Println("Starting server on http://localhost:3000/")
    err := http.ListenAndServe(":3000", nil)
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

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        if r.URL.Path != "/ascii-art" {
            http.Error(w, "400 - Bad Request", 400)
            return
        }
        text := r.FormValue("text")
        banner := r.FormValue("banner")
        if len(text) == 0 || len(banner) == 0 {
            http.Error(w, "400 - Bad Request", http.StatusBadRequest)
        }

        asciiArt, errs := AsciiArtMaker(text, banner)
        tmpl, err := template.ParseFiles("template.html")
        errs = append(errs, err)

        //Handing template err and AsciiConverter errs
        for i := range errs {
            if errs[i] != nil {
                http.Error(w, "404 - Not Found", http.StatusNotFound)
                return
            }
        }

        data := Data{
            Text:     text,
            Banner:   banner,
            AsciiArt: asciiArt,
        }
        tmpl.Execute(w, data)
    } 
    if r.Method == "GET" {
        if r.URL.Path != "/" {
            http.Error(w, "400 - Bad Request", 400)
            return
        }
        //http.ServeFile(w, r, "template.html")
        tmpl, err := template.ParseFiles("template.html")
        if err != nil {
            http.Error(w, "404 - Not Found", 404)
            return
        }
        data := Data{
            Text:     "",
            AsciiArt: "",
        }
        tmpl.ExecuteTemplate(w, "template.html", data)
    }
}
