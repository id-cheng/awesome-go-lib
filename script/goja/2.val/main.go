package main

import (
	"fmt"
	"github.com/dop251/goja"
)

func main() {
	newField()
	newMap()
	newArray()
}

func newField() {
	type S struct {
		Field int
	}
	vm := goja.New()
	a := []S{{1}, {2}} // slice of literal structs
	vm.Set("a", &a)
	vm.RunString(`
    let tmp = {Field: 1};
    a[0] = tmp;
    a[1] = tmp;
    tmp.Field = 2;
`)
	fmt.Println(a)

	a1 := []any{S{1}, S{2}}
	vm.Set("a1", &a1)
	vm.RunString(`
	   a1[0].Field === 1; // true
	   a1[0].Field = 2;
	   a1[0].Field === 2; // FALSE, because what it really did was copy a1[0] set its Field to 2 and immediately drop it
`)
}

func newMap() {
	vm := goja.New()
	m := map[string]any{}
	vm.Set("m", m)
	vm.RunString(`
	var obj = {test: false};
	m.obj = obj;
	obj.test = true;
`)
	fmt.Println(m["obj"])
	fmt.Println(m["obj"].(map[string]any)["test"])
}

// NewArray 不可寻址的结构、切片和数组被复制
func newArray() {
	vm := goja.New()
	a := []any{1}
	vm.Set("a", a)
	vm.RunString(`a.push(2); a[0] = 0;`)
	fmt.Println(a[0]) // prints "1"
}
