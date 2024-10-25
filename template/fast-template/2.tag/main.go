package main

import (
	"fmt"
	"github.com/valyala/fasttemplate"
	"io"
)

func main() {
	template := `name: {{name}} age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "name":
			return w.Write([]byte("cheng"))
		case "age":
			return w.Write([]byte("18"))
		default:
			return 0, nil
		}
	})
	fmt.Println(s)

	template = `name: {{name}} age: {{age}}`
	t = fasttemplate.New(template, "{{", "}}")
	m := map[string]interface{}{"name": "cheng"}
	s1 := t.ExecuteString(m)
	fmt.Println(s1)

	s2 := t.ExecuteStringStd(m)
	fmt.Println(s2)
}
