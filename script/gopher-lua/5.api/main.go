package main

import (
	"github.com/yuin/gopher-lua"
)

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}

func main() {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("double", L.NewFunction(Double))
	L.GetGlobal("double")
	if err := L.DoString(`print(double(10))`); err != nil {
		panic(err)
	}
}
