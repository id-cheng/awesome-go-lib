package main

import lua "github.com/yuin/gopher-lua"

func main() {
	L := lua.NewState(lua.Options{
		CallStackSize:       120,  // this is the maximum callstack size of this LState
		MinimizeStackMemory: true, // Defaults to `false` if not specified. If set, the callstack will auto grow and shrink as needed up to a max of `CallStackSize`. If not set, the callstack will be fixed at `CallStackSize`.
	})
	defer L.Close()
}
