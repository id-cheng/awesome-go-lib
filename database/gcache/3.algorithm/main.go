package main

import (
	"fmt"
	"github.com/bluele/gcache"
)

func main() {
	lfu()
	lru()
	arc()
	simple()
}

func simple() {
	// size: 10
	gc := gcache.New(10).Build()
	gc.Set("key", "value")
	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func arc() {
	// size: 10
	gc := gcache.New(10).
		ARC().
		Build()
	gc.Set("key", "value")
}

func lru() {
	// size: 10
	gc := gcache.New(10).
		LRU().
		Build()
	gc.Set("key", "value")
}

func lfu() {
	// size: 10
	// Least-Frequently Used (LFU)
	gc := gcache.New(10).
		LFU().
		Build()
	gc.Set("key", "value")
}
