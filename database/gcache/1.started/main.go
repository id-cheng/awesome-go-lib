package main

import (
	"fmt"
	"github.com/bluele/gcache"
	"time"
)

func main() {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.Set("key1", "val1")
	gc.SetWithExpire("key2", "val2", time.Second*10)
	value, _ := gc.Get("key1")
	fmt.Println("Get:", value)

	// Wait for value to expire
	time.Sleep(time.Second * 10)

	value, err := gc.Get("key2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Get:", value)
}
