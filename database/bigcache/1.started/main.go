package main

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

var ctx = context.Background()
var conf = bigcache.DefaultConfig(10 * time.Minute)

func main() {
	cache, _ := bigcache.New(ctx, conf)
	cache.Set("my-unique-key", []byte("value"))
	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}
