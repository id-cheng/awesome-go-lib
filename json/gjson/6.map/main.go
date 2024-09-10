package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"jia","last":"cheng"},"age":20}`

func main() {
	m, ok := gjson.Parse(json).Value().(map[string]any)
	if ok {
		fmt.Println(m["name"])
		fmt.Println(m["age"])
	}
}
