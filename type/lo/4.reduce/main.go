package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	fmt.Println(sum)

	result := lo.ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})
	fmt.Println(result)
}
