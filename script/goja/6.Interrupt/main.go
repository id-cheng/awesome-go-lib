package main

import (
	"fmt"
	"github.com/dop251/goja"
	"time"
)

func main() {
	const SCRIPT = `
    var i = 0;
    for (;;) {
        i++;
    }
    `

	vm := goja.New()
	time.AfterFunc(200*time.Millisecond, func() {
		vm.Interrupt("halt")
	})

	_, err := vm.RunString(SCRIPT)
	fmt.Println(err)
	// err is of type *InterruptError and its Value() method returns whatever has been passed to vm.Interrupt()
}
