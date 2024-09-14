package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	const SCRIPT = `
	function sum(a, b) {
		return a + b;
	}
`
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}
	sum, ok := goja.AssertFunction(vm.Get("sum"))
	if !ok {
		panic("Not a function")
	}

	res, err := sum(goja.Undefined(), vm.ToValue(10), vm.ToValue(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	var Sum func(int, int) int
	err = vm.ExportTo(vm.Get("sum"), &Sum)
	if err != nil {
		panic(err)
	}
	fmt.Println(Sum(10, 2))
}
