package main

import (
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	process(c)
}

func process(stream chan int) {
	fmt.Println("process")
	p := pool.New().WithMaxGoroutines(10)
	for elem := range stream {
		e := elem
		p.Go(func() {
			fmt.Println(e)
		})
	}
	p.Wait()
}
