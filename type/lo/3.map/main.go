package main

import (
	"fmt"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func main() {
	strs := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Println(strs)

	matching := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return x, true
		}
		return "", false
	})
	fmt.Println(matching)

	flatMap := lo.FlatMap([]int64{0, 1, 2}, func(x int64, _ int) []string {
		return []string{
			strconv.FormatInt(x, 10),
			strconv.FormatInt(x, 10),
		}
	})
	fmt.Println(flatMap)
}
