package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

// 监听文件
func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err = watcher.Add("./"); err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}
