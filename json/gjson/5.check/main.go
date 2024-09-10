package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"jia","last":"cheng"},"age":20}`

func main() {
	value := gjson.Get(json, "name.last")
	if !value.Exists() {
		println("no last name")
	} else {
		println(value.String())
	}

	if !gjson.Valid(json) {
		fmt.Println("invalid json")
	}
}
