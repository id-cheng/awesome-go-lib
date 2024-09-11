package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]any{
		"name":  "cheng",
		"age":   18,
		"email": "foo@bar.com",
	}

	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md mapstructure.Metadata
	var result Person
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	if err := decoder.Decode(input); err != nil {
		panic(err)
	}

	fmt.Println("keys:", md.Keys)
	fmt.Println("Unused keys:", md.Unused)
}
