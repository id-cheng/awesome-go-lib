package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	intSlice := []int{1, 2, 3}
	intArray := [3]int{1, 10, 100}
	// ToIntSlice
	fmt.Println(cast.ToIntSlice(intSlice))
	fmt.Println(cast.ToIntSlice(intArray))

	anySlice := []any{1, 10.0, "go"}
	strSlice := []string{"Java", "go", "python"}
	str := " rust  go  c++ "
	// ToStringSlice
	fmt.Println(cast.ToStringSlice(intSlice))
	fmt.Println(cast.ToStringSlice(anySlice))
	fmt.Println(cast.ToStringSlice(strSlice))
	fmt.Println(cast.ToStringSlice(str))
}
