package main

import (
	"errors"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	//Just创建一个仅有若干固定数据的Observable
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}

	observable = rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	<-observable.ForEach(func(v any) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})

	observable = rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	ch = observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		}
	}
}
