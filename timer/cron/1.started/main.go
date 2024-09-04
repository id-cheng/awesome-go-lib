package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	done := make(<-chan struct{})
	c := cron.New()
	var count int
	c.AddFunc("@every 1s", func() {
		fmt.Println("hello ", count)
		count++
	})
	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
	})

	c.AddFunc("@daily", func() {
		fmt.Println("Every day")
	})
	c.AddFunc("@weekly", func() {
		fmt.Println("Every week")
	})
	c.Start()
	<-done
}
