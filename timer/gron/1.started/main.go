package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"time"
)

func main() {
	done := make(<-chan bool)
	cron := gron.New()
	var count int
	cron.AddFunc(gron.Every(time.Second), func() {
		fmt.Printf("helloWorld %d\n", count)
		count++
	})
	cron.Start()
	<-done
}
