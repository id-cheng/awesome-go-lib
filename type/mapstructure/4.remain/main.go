package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name  string
	Age   int
	Other map[string]any `mapstructure:",remain"`
}

func main() {
	input := map[string]any{
		"name":  "cheng",
		"age":   20,
		"email": "mit@example.com",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}
