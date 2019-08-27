package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	PORT  = ":8080"
	INDEX = "templates/index.html"
)

var GlobalStore = map[string]int{}

func main() {

	go checker()

	if len(os.Args) == 2 {
		cliGetGist()
	}

	log.Printf("Started Github Public Gist Query Application on port [%v] ", PORT)

	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)

	router.HandleFunc("/query", QueryHandler)

	router.HandleFunc("/api", ApiHandler).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(PORT, nil))

}

// retrive the public gists for a user passed as a commandline argument

func cliGetGist() {
	user := os.Args[1]

	g, err := GetGist(user)

	if err != nil {
		log.Println(err)
	}

	if _, exists := GlobalStore[user]; !exists {
		GlobalStore[user] = g.Counter

	}
}
