package main

import (
	"fmt"
	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(20).
		LRU().
		LoaderFunc(func(key any) (any, error) {
			return "ok", nil
		}).
		Build()
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
