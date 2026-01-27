package main

import "net/http"

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Unit Converter"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	http.ListenAndServe("localhost:8080", mux)
}
