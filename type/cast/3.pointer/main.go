package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	p := new(int)
	*p = 10
	fmt.Println(cast.ToInt(p))

	pp := &p
	fmt.Println(cast.ToInt(pp))
}
