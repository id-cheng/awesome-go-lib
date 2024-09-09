package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func main() {
	output, _ := expr.Eval("greet + name", map[string]any{
		"greet": "Hello ",
		"name":  "world!",
	})
	fmt.Println(output)
}
