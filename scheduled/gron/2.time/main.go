package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"time"
)

func main() {
	done := make(<-chan bool)
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Second), func() {
		fmt.Println("run every second")
	})
	c.AddFunc(gron.Every(1*time.Minute), func() {
		fmt.Println("run every minute")
	})
	c.AddFunc(gron.Every(1*time.Hour), func() {
		fmt.Println("run every hour")
	})

	t, _ := time.ParseDuration("10m10s")
	c.AddFunc(gron.Every(t), func() {
		fmt.Println("run 10 minutes 10 seconds.")
	})
	c.Start()
	<-done
}
