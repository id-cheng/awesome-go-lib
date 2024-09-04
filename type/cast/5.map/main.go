package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	m1 := map[string]string{
		"name": "go",
		"age":  "18",
	}

	m2 := map[string]any{
		"name": "Java",
		"age":  30,
	}

	m3 := map[any]string{
		"name": "c",
		"age":  "48",
	}

	jsonStr := `{"name":"rust", "age":"9"}`
	fmt.Println(cast.ToStringMapString(m1))
	fmt.Println(cast.ToStringMapString(m2))
	fmt.Println(cast.ToStringMapString(m3))
	fmt.Println(cast.ToStringMapString(jsonStr))
	fmt.Println(cast.ToStringMapInt(map[string]any{"age": "18"}))
	fmt.Println(cast.ToStringMapInt(`{"age":9}`))
	fmt.Println(cast.ToStringMapBool(map[string]any{"is_deleted": "true"}))
	fmt.Println(cast.ToStringMapBool(`{"is_deleted":false}`))
}
