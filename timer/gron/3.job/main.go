package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"time"
)

type Job struct {
	Name string
}

func (j Job) Run() {
	fmt.Printf("hello %s %s \n", j.Name, time.Now().Format(time.DateTime))
}

func main() {
	done := make(<-chan bool)
	j1 := Job{Name: "j1"}
	j2 := Job{Name: "j2"}
	cron := gron.New()
	cron.Add(gron.Every(time.Second), j1)
	cron.Add(gron.Every(time.Minute), j2)
	cron.Start()
	<-done
}
