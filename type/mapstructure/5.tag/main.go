package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string `mapstructure:"person_name"`
	Age  int    `mapstructure:"person_age"`
}

func main() {

	input := map[string]interface{}{
		"person_name": "cheng",
		"person_age":  19,
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)

}
