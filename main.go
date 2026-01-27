package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func convertHandler(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	//Fetch the form values from the name field in html tag and validate them before passing  them to the function
	value, err := strconv.ParseFloat(strings.TrimSpace(r.FormValue("value")), 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	toUnit, err := stringToUnit(strings.TrimSpace(r.FormValue("to")))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	fromUnit, err := stringToUnit(strings.TrimSpace(r.FormValue("from")))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	convertedValue, err := convertLengthUnit(value, fromUnit, toUnit)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(rw, "Result: %f", convertedValue)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /convert", convertHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
