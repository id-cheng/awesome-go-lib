package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Hello %s\n", call.Argument(0).String())
		return otto.Value{}
	})
	vm.Set("sum", func(call otto.FunctionCall) otto.Value {
		a, _ := call.Argument(0).ToInteger()
		b, _ := call.Argument(1).ToInteger()
		result, _ := vm.ToValue(a + b)
		return result
	})
	vm.Run(`sayHello('cheng')
                result = sum(1,2)
				console.log('sum:'+result)`)
}
