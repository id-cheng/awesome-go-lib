package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Job struct {
	Name string
}

func (j Job) Run() {
	fmt.Printf("hello %s %s \n", j.Name, time.Now().Format(time.DateTime))
}

func main() {
	done := make(<-chan struct{})
	c := cron.New()
	id, err := c.AddJob("@every 1s", Job{"j1"})
	if err != nil {
		return
	}
	fmt.Println("job_id", id)
	c.Start()
	<-done
}
