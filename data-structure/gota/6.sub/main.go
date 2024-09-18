package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"strings"
)

func main() {
	jsonStr := `[{"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`
	df := dataframe.ReadJSON(strings.NewReader(jsonStr))
	sub := df.Subset([]int{1, 2})
	frame := sub.Select([]int{0, 1})
	fmt.Println(sub)
	fmt.Println(frame)
}
