package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	done := make(<-chan struct{})
	location, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(location))
	var count int
	id, err := c.AddFunc("* * * * *", func() {
		fmt.Println("hello ", count)
		count++
	})
	if err != nil {
		return
	}
	fmt.Println("id:", id)
	c.Start()
	<-done
}
