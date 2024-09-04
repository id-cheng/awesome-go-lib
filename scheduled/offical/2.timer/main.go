package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// 一次性定时器timer
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	go func() {
		select {
		case <-timer.C:
			defer wg.Done()
			log.Println("timer1")
		}
	}()
	// 一次性定时器timer
	time.AfterFunc(time.Second, func() {
		defer wg.Done()
		log.Println("timer2")
	})
	wg.Wait()
}
