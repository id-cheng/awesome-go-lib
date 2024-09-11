package main

import (
	"fmt"
	"github.com/sourcegraph/conc/iter"
)

func main() {
	iterForeach([]int{1, 2, 3})
	ints := iterMap([]int{1, 2, 3}, func(i *int) int {
		return *i
	})
	fmt.Println(ints)
}
func iterMap(input []int, f func(*int) int) []int {
	fmt.Println("iterMap")
	return iter.Map(input, f)
}

func iterForeach(values []int) {
	fmt.Println("iterForeach")
	iter.ForEach(values, func(i *int) {
		fmt.Println(*i)
	})
}
