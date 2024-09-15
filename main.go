package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//Register the handler to a root
	http.HandleFunc("/", htmlHandler)

	//Start the web Server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server :", err)
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := Read("template.html")
	if err != nil {
		http.Error(w, "internal Server Error", http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "text/html")
	w.Write(htmlContent)
}

func Read(fileName string) ([]byte, error) {
	
	//Open File.
    file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening file :", fileName)
		return []byte{}, err
	}

	defer file.Close()
	
	//Get file info.
    fileInfo, err := file.Stat()
    if err != nil {
        log.Println("Error getting file stats:", err)
		return []byte{}, err
    }

    //Get file size.
    fileSize := fileInfo.Size()
    data := make([]byte, fileSize)

	//Reading data.
    _, err = file.Read(data)
    if err != nil {
		log.Println("Error reading the file:", err)
		return []byte{}, err
    }
	return data, nil
}
