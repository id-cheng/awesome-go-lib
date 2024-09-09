package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	lv := L.Get(-1) // get the value at the top of the stack
	if str, ok := lv.(lua.LString); ok {
		// lv is LString
		fmt.Println(string(str))
	}
	if lv.Type() != lua.LTString {
		fmt.Println("not str")
	}

	if tbl, ok := lv.(*lua.LTable); ok {
		// lv is LTable
		fmt.Println(L.ObjLen(tbl))
	}
	fmt.Println(lv.Type())
}
