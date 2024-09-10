package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	jsonStr := `
	{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "age":20,
      "enable":true,
	  "children": ["Sara","Alex","Jack"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	  ]
	}`
	children := gjson.Get(jsonStr, "children")
	fmt.Println(children.Type)

	num := gjson.Get(jsonStr, "children.#")
	fmt.Println(num.Type)

	child := gjson.Get(jsonStr, "children.1")
	fmt.Println(child.Type)

	age := gjson.Get(jsonStr, "age")
	fmt.Println(age.Type)

	enable := gjson.Get(jsonStr, "enable")
	fmt.Println(enable.Type)
}
