package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type User struct {
	Name     string
	Age      int
	Accuracy float64
	ignored  bool // ignored since unexported
}

func main() {
	users := []User{
		{"Aram", 17, 0.2, true},
		{"Juan", 18, 0.8, true},
		{"Ana", 22, 0.5, true},
	}
	df := dataframe.LoadStructs(users)
	fmt.Println(df.String())

	df = dataframe.LoadRecords(
		[][]string{
			{"A", "B", "C", "D"},
			{"a", "4", "5.1", "true"},
			{"k", "5", "7.0", "true"},
			{"k", "4", "6.0", "true"},
			{"a", "2", "7.1", "false"},
		},
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.Float),
		dataframe.WithTypes(map[string]series.Type{
			"A": series.String,
			"D": series.Bool,
		}),
	)
	fmt.Println(df.String())
}
