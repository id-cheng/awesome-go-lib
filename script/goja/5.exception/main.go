package main

import (
	"errors"
	"fmt"
	"github.com/dop251/goja"
)

var runtime *goja.Runtime

func exception() {
	panic(runtime.ToValue("Error"))
}
func main() {
	vm := goja.New()
	_, err := vm.RunString(`throw("Test")`)

	jserr := new(goja.Exception)
	if errors.As(err, &jserr) {
		if jserr.Value().Export() != "Test" {
			panic("wrong value")
		} else {
			fmt.Println("Test")
		}
	}

	vm.Set("exception", exception)
	_, err = vm.RunString(`
		try {
			exception();
		} catch(e) {
			if (e === "Error") {
				throw e;
			}
		}
    `)
	if err != nil {
		panic(err)
	}
}
