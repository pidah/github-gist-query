package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	PORT  = ":8080"
	INDEX = "templates/index.html"
)

var GlobalStore = map[string]int{}

func main() {

	log.Printf("Started Github Public Gist Query Application on port [%v] ", PORT)

	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)

	router.HandleFunc("/query", QueryHandler)

	router.HandleFunc("/json", JsonHandler).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(PORT, nil))

}
