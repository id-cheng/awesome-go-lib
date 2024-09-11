package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
)

func main() {
	var wg conc.WaitGroup
	defer wg.Wait()
	start(&wg)
	iter(&wg)
	doPanic(&wg)
}

func start(wg *conc.WaitGroup) {
	wg.Go(func() {
		fmt.Println("helloWorld")
	})
}
func iter(wg *conc.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			fmt.Println("iter:", i)
		})
	}
}

func doPanic(wg *conc.WaitGroup) {
	wg.Go(func() {
		panic("panic")
	})
}
