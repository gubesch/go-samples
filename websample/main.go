package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)




func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets")

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}


func main() {
	router := newRouter()

	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Devs!")
}