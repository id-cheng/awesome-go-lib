package main

import (
	"github.com/samber/lo"
)

func main() {
	lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})

}
