package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github" // with go modules disabled
	"log"
)

type Gists struct {
	Gist    []string `json:"gist"`
	Counter int
}

func GetGist(githubUser string) (Gists, error) {
	if githubUser == "" {
		return Gists{}, nil
	}

	client := github.NewClient(nil)

	// list all public gists for githubUser
	var list []string

	gists, _, err := client.Gists.List(context.Background(), githubUser, nil)

	if err != nil {
		log.Println(err)
	}

	for _, g := range gists {
		for _, v := range g.Files {
			list = append(list, *v.RawURL)
		}
	}
	gistCount := len(list)
	G := Gists{list, gistCount}
	jsonList, err := json.MarshalIndent(list, "", "    ")

	log.Printf("public gists for %s: %s\n", githubUser, string(jsonList))

	return G, nil
}
