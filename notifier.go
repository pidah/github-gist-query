package main

import (
	"github.com/0xAX/notificator"
	"log"
)

func notify(user string) error {
	var notify *notificator.Notificator
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Github Gist App",
	})

	err := notify.Push("New Public Gist found for "+user,
		"refresh browser or submit new query for "+user,
		"static/image/octocat.png",
		notificator.UR_NORMAL)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
