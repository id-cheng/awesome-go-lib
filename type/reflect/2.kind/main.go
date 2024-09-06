package main

import (
	"fmt"
	"reflect"
)

type Int int

// 静态类型分别为int和Int是不同的
// 但是它们的种类都是int
func main() {
	var i int
	var j Int

	i = int(j) // 强制转换

	ti := reflect.TypeOf(i)
	fmt.Println("type of i:", ti.String())

	tj := reflect.TypeOf(j)
	fmt.Println("type of j:", tj.String())

	fmt.Println("kind of i:", ti.Kind())
	fmt.Println("kind of j:", tj.Kind())
}
