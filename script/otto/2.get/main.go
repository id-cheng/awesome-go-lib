package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Run(`sum = 1 + 1
    console.log("sum:" + sum)`)

	val, _ := vm.Get("sum")
	valInt, _ := val.ToInteger()
	fmt.Println("val:", valInt)
}
