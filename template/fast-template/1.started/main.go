package main

import (
	"fmt"

	"github.com/valyala/fasttemplate"
)

func main() {
	template := `name: {{name}},age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	s1 := t.ExecuteString(map[string]any{
		"name": "xie",
		"age":  "18",
	})
	s2 := t.ExecuteString(map[string]any{
		"name": "cheng",
		"age":  "20",
	})
	fmt.Println(s1)
	fmt.Println(s2)

	template = ` name: [name] age: [age] `
	s := fasttemplate.ExecuteString(template, "[", "]", map[string]any{
		"name": "cheng",
		"age":  "18",
	})
	fmt.Println(s)
}
