package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	fil := df.Filter(
		dataframe.F{Colname: "A", Comparator: series.Eq, Comparando: "a"},
		dataframe.F{Colname: "B", Comparator: series.Greater, Comparando: 1},
	)
	fmt.Println(fil.String())
	filAlt := df.FilterAggregation(
		dataframe.And,
		dataframe.F{Colname: "A", Comparator: series.Eq, Comparando: "a"},
		dataframe.F{Colname: "B", Comparator: series.Greater, Comparando: 1},
	)
	fmt.Println(filAlt.String())

	filAlt = df.FilterAggregation(
		dataframe.Or,
		dataframe.F{Colname: "A", Comparator: series.Eq, Comparando: "a"},
		dataframe.F{Colname: "B", Comparator: series.Greater, Comparando: 1},
	)
	fmt.Println(filAlt.String())
}

var df = dataframe.LoadMaps(
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
