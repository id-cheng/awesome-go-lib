package main

import (
	"log"
	"time"
)

func main() {
	// 周期性定时器ticker
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var count int
	for range ticker.C {
		log.Println("count:", count)
		count++
	}
}
