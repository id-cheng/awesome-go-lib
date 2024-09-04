package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"sync/atomic"
	"time"
)

type panicJob struct {
	arrays [5]int
	count  int
}

func (j *panicJob) Run() {
	log.Println("panic job ", j.count)
	j.count++
	j.arrays[j.count] = j.count
}

type delayJob struct {
	count int
}

func (d *delayJob) Run() {
	time.Sleep(2 * time.Second)
	log.Println("delay job ", d.count)
	d.count++
}

type skipJob struct {
	count int32
}

func (s *skipJob) Run() {
	atomic.AddInt32(&s.count, 1)
	log.Println("skip job ", s.count)
	if atomic.LoadInt32(&s.count) == 1 {
		time.Sleep(2 * time.Second)
	}
}

func main() {
	done := make(<-chan struct{})
	c := cron.New()
	j1, _ := c.AddJob("@every 1s", cron.NewChain(
		cron.Recover(cron.DiscardLogger)).Then(&panicJob{}))
	j2, _ := c.AddJob("@every 1s", cron.NewChain(
		cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&delayJob{}))
	j3, _ := c.AddJob("@every 1s", cron.NewChain(
		cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&skipJob{}))
	fmt.Println("j1:", j1)
	fmt.Println("j2:", j2)
	fmt.Println("j3:", j3)
	c.Start()
	<-done
}
