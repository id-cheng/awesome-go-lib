package main

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"time"
)

func main() {
	// create a scheduler
	s, _ := gocron.NewScheduler()

	// add a job to the scheduler
	j, _ := s.NewJob(
		gocron.DurationJob(
			1*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println(a, b)
			},
			"hello", 1,
		),
	)
	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute):
	}

	// when you're done, shut it down
	_ = s.Shutdown()
}
