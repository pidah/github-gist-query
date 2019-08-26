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
	User string
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

	if _, ok := GlobalStore[user]; !ok {
		GlobalStore[user] = g.Counter
	}

	if err != nil {
		log.Printf(err.Error())
	}

}

func JsonHandler(w http.ResponseWriter, r *http.Request) {

	accept := r.Header.Get("Accept")

	switch accept {

	case "application/json":

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
		jsonData, _ := json.Marshal(g)

		fmt.Fprint(w, bytes.NewBuffer(jsonData))
		return
	}
}
