package main

import (
	"fmt"
	"github.com/expr-lang/expr"
)

type Env struct {
	X int
	Y int
}

func envMap() {
	env := map[string]any{
		"foo": 100,
		"bar": 200,
	}
	program, _ := expr.Compile(`foo + bar`, expr.Env(env))
	output, _ := expr.Run(program, env)
	fmt.Println(output)

	env = map[string]any{
		"greet":   "Hello, %v!",
		"names":   []string{"world", "you"},
		"sprintf": fmt.Sprintf,
	}
	program, _ = expr.Compile(`sprintf(greet, names[0])`, expr.Env(env))
	output, _ = expr.Run(program, env)
	fmt.Println(output)
}

func envStruct() {
	program, _ := expr.Compile(`X + Y`, expr.Env(Env{}))
	output, _ := expr.Run(program, Env{1, 2})
	fmt.Println(output)
}
func main() {
	envMap()
	envStruct()
}
