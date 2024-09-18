package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	df := dataframe.New()
	df2 := df.Set(
		[]int{0, 2},
		dataframe.LoadRecords(
			[][]string{
				{"A", "B", "C", "D"},
				{"b", "4", "6.0", "true"},
				{"c", "3", "6.0", "false"},
			},
		),
	)
	fmt.Println(df2)
}
