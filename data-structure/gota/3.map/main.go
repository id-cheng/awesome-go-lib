package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	df := dataframe.LoadMaps(
		[]map[string]any{
			{
				"A": "a",
				"B": 1,
				"C": true,
				"D": 0,
			},
			{
				"A": "b",
				"B": 2,
				"C": true,
				"D": 0.5,
			},
		},
	)
	fmt.Println(df.String())
}
