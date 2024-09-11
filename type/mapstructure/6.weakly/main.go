package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	input := map[string]any{
		"name":   123,              // number => string
		"age":    "18",             // string => number
		"emails": map[string]any{}, // empty map => empty array
	}

	var result Person
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}
