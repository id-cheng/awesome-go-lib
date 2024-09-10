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
	fmt.Println(children.Array())

	num := gjson.Get(jsonStr, "children.#")
	fmt.Println(num.Int())

	child := gjson.Get(jsonStr, "children.1")
	fmt.Println(child.String())

	child = gjson.Get(jsonStr, "child*.2")
	fmt.Println(child.String())

	child = gjson.Get(jsonStr, "c?ildren.0")
	fmt.Println(child.String())

	friends := gjson.Get(jsonStr, "friends.#.first")
	fmt.Println(friends.IsArray())

	friend := gjson.Get(jsonStr, "friends.1.last")
	fmt.Println(friend.String())

	enable := gjson.Get(jsonStr, "enable")
	fmt.Println(enable.Bool())
}
