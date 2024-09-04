package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// ToString
	fmt.Println(cast.ToString(10))
	fmt.Println(cast.ToString(10.10))
	fmt.Println(cast.ToString([]byte("cast")))
	fmt.Println(cast.ToString(nil))

	var foo any = "any type"
	fmt.Println(cast.ToString(foo)) // one more time

	// ToInt
	fmt.Println(cast.ToInt(10))
	fmt.Println(cast.ToInt(10.10))
	fmt.Println(cast.ToInt("10"))
	fmt.Println(cast.ToInt(true))
	fmt.Println(cast.ToInt(false))

	var ten any = 10
	fmt.Println(cast.ToInt(ten))
	fmt.Println(cast.ToInt(nil))
}
