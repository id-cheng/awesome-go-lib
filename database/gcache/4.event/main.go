package main

import (
	"fmt"
	"github.com/bluele/gcache"
)

func main() {
	evicted()
	added()
}

func added() {
	gc := gcache.New(2).
		AddedFunc(func(key, value interface{}) {
			fmt.Println("added key:", key)
		}).
		Build()
	for i := 0; i < 3; i++ {
		gc.Set(i, i*i)
	}
}

func evicted() {
	gc := gcache.New(2).
		EvictedFunc(func(key, value any) {
			fmt.Println("evicted key:", key)
		}).
		Build()
	for i := 0; i < 3; i++ {
		gc.Set(i, i*i)
	}
}
