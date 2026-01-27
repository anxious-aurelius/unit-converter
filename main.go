package main

import (
	"html/template"
	"log"
	"net/http"
)

type formData struct {
	Units []string
}

func homeHandler(rw http.ResponseWriter, _ *http.Request) {

	lengthData := formData{unitToString()}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	files, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	files.Execute(rw, lengthData)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
