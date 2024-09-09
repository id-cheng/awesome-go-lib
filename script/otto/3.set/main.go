package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Set("val", 10)
	vm.Set("str", "hello")
	vm.Run(`
    console.log("val:" + val)
    console.log("str:" + str)`)

	str, _ := vm.Run("str.length")
	{
		length, _ := str.ToInteger()
		fmt.Println("length:", length)
	}
}
