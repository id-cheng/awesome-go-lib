package main

import (
	"fmt"
	"github.com/sourcegraph/conc/stream"
)

func main() {
	in := make(chan int, 10)
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
	out := make(chan int, 10)
	mapStream(in, out, func(i int) int {
		return i * 2
	})
	close(out)
	for i := range out {
		fmt.Println(i)
	}
}
func mapStream(
	in chan int,
	out chan int,
	f func(int) int,
) {
	s := stream.New().WithMaxGoroutines(10)
	for elem := range in {
		e := elem
		s.Go(func() stream.Callback {
			res := f(e)
			return func() {
				out <- res
			}
		})
	}
	s.Wait()
}
