package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tmpl *template.Template

type formData struct {
	Units      []string
	ShowResult bool
	Result     float64
}

func homeHandler(rw http.ResponseWriter, _ *http.Request) {
	data := formData{unitsToString(), false, 0}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(rw, data)
}

func convertHandler(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	//Fetch the form values from the name field in HTML tag and validate them before passing  them to the function
	value, err := strconv.ParseFloat(strings.TrimSpace(r.FormValue("value")), 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	toUnit, err := stringToUnit(strings.TrimSpace(r.FormValue("to")))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	fromUnit, err := stringToUnit(strings.TrimSpace(r.FormValue("from")))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	convertedValue, err := convertLengthUnit(value, fromUnit, toUnit)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	data := formData{unitsToString(), true, convertedValue}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(rw, data)
}

func main() {
	// parsing template files in the package level variable declared
	var err error
	tmpl, err = template.ParseFiles("./templates/home.html")
	if err != nil {
		log.Fatal("Error loading template files. Quitting....")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /convert", convertHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
