package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type gist_json struct {
	User string `json:"user"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("RootHandler started and redirecting to the index page")

	t, _ := template.ParseFiles(INDEX)

	t.Execute(w, nil)

}

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	//Retrieve the HTML form parameter of POST method
	user := r.FormValue("github-user")

	g, err := GetGist(user)

	if err != nil {
		log.Printf("GetGistError: %s, %v",
			err.Error(), http.StatusInternalServerError)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, _ := template.ParseFiles(INDEX)

	t.Execute(w, g)

	if _, exists := GlobalStore[user]; !exists {
		GlobalStore[user] = g.Counter
	}
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var request gist_json

	err := decoder.Decode(&request)

	g, err := GetGist(request.User)

	if err != nil {
		log.Printf("GetGistError: %s, %v",
			err.Error(), http.StatusInternalServerError)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, exists := GlobalStore[request.User]; !exists {
		GlobalStore[request.User] = g.Counter
	}

	jsonData, _ := json.Marshal(g.Gist)

	fmt.Fprint(w, bytes.NewBuffer(jsonData))
	return
}
