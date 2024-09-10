package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	var json = []byte(`{"name":{"first":"jia","last":"cheng"},"age":20`)
	result := gjson.GetBytes(json, "name.first")
	fmt.Println(result.Str)
}
