package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/expr-lang/expr"
)

func allowUndefined() {
	code := `name == nil ? "Hello, world!" : sprintf("Hello, %v!", name)`
	//code := `name == "" ? "Hello, world!" : sprintf("Hello, %v!", name)`

	env := map[string]any{
		"sprintf": fmt.Sprintf,
	}

	options := []expr.Option{
		expr.Env(env),
		expr.AllowUndefinedVariables(), // Allow to use undefined variables.
	}

	program, _ := expr.Compile(code, options...)
	output, _ := expr.Run(program, env)
	fmt.Printf("%v\n", output)

	env["name"] = "you" // Define variables later on.
	output, _ = expr.Run(program, env)
	fmt.Printf("%v\n", output)
}

func asBool() {
	env := map[string]int{"foo": 0}
	program, _ := expr.Compile("foo >= 0", expr.Env(env), expr.AsBool())
	output, _ := expr.Run(program, env)
	fmt.Printf("%v", output.(bool))
}

func asFloat64() {
	program, _ := expr.Compile("10", expr.AsFloat64())
	output, _ := expr.Run(program, nil)
	fmt.Printf("%v", output.(float64))
}

func asInt() {
	program, _ := expr.Compile("12", expr.AsInt())
	output, _ := expr.Run(program, nil)
	fmt.Printf("%T(%v)\n", output, output)
}

func asInt64() {
	env := map[string]any{
		"rating": 5.5,
	}
	program, _ := expr.Compile("rating", expr.Env(env), expr.AsInt64())
	output, _ := expr.Run(program, env)
	fmt.Printf("%v\n", output.(int64))
}

func asKind() {
	program, _ := expr.Compile("{a: 1, b: 2}", expr.AsKind(reflect.Map))
	output, _ := expr.Run(program, nil)
	fmt.Printf("%v\n", output)
}

func asOperator() {
	code := `
		Now() > CreatedAt &&
		(Now() - CreatedAt).Hours() > 24
	`

	type Env struct {
		CreatedAt time.Time
		Now       func() time.Time
		Sub       func(a, b time.Time) time.Duration
		After     func(a, b time.Time) bool
	}

	options := []expr.Option{
		expr.Env(Env{}),
		expr.Operator(">", "After"),
		expr.Operator("-", "Sub"),
	}

	program, _ := expr.Compile(code, options...)
	env := Env{
		CreatedAt: time.Date(2018, 7, 14, 0, 0, 0, 0, time.UTC),
		Now:       func() time.Time { return time.Now() },
		Sub:       func(a, b time.Time) time.Duration { return a.Sub(b) },
		After:     func(a, b time.Time) bool { return a.After(b) },
	}

	output, _ := expr.Run(program, env)
	fmt.Printf("%v\n", output)
}

func warnOnAny() {
	_, err := expr.Compile(`[42, true, "yes"][0]`, expr.AsInt(), expr.WarnOnAny())
	fmt.Printf("%v\n", err)
}

func main() {
	allowUndefined()
	asBool()
	asFloat64()
	asInt()
	asInt64()
	asKind()
	asOperator()
	warnOnAny()
}
