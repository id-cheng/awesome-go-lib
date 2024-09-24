package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
)

func main() {
	m := make(map[int]bool)
	m[1] = true
	m[2] = false

	fmt.Println(linq.From(m).Results())

	input := []linq.KeyValue{
		{10, true},
	}
	m1 := make(map[int]bool)
	linq.From(input).ToMap(&m1)
	fmt.Println(m1)
}
