package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	wat, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}

	defer wat.Close()

	go func() {
		for {
			select {
			case event, ok := <-wat.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				switch event.Op {
				case fsnotify.Create:
					fmt.Println("created file:", event.Name)
				case fsnotify.Remove:
					fmt.Println("removed file:", event.Name)
				}
			case err, ok := <-wat.Errors:
				if !ok {
					return
				}

				log.Println("error:", err)

			}
		}
	}()

	if err := wat.Add("./"); err != nil {
		log.Fatalln(err)
	}

	<-make(chan bool)
}
