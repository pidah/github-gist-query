package main

import (
	"log"
	"time"
)

func checker() {

	for {
		for user, currentCounter := range GlobalStore {
			new, _ := GetGist(user)
			if new.Counter > currentCounter {
				log.Printf("New Public gist found for %s !", user)
				notify(user)
				GlobalStore[user] = new.Counter
			}
		}
		time.Sleep(10 * time.Second)
	}

}
