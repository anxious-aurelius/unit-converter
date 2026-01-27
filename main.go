package main

import "net/http"

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	files.Execute(rw, nil)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	http.ListenAndServe("localhost:8080", mux)
}
