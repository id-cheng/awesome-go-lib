package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
)

func main() {
	times := lo.Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Println(times)
}
