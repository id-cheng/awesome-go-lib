package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Family struct {
	LastName string
}
type Location struct {
	City string
}
type Person struct {
	*Family   `mapstructure:",omitempty"`
	*Location `mapstructure:",omitempty"`
	Age       int
	FirstName string
}

func main() {
	result := map[string]any{}
	input := Person{FirstName: "Somebody"}
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", result)
}
