package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

func main() {
	input := map[string]any{
		"name":   "cheng",
		"age":    18,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mit",
		},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
