package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	type S struct {
		Field int `json:"field"`
	}
	vm.Set("s", S{Field: 42})
	res, _ := vm.RunString(`s.field`) // without the mapper it would have been s.Field
	fmt.Println(res.Export())
}
