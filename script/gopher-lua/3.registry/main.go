package main

import lua "github.com/yuin/gopher-lua"

func main() {
	L := lua.NewState(lua.Options{
		RegistrySize:     1024 * 20, // this is the initial size of the registry
		RegistryMaxSize:  1024 * 80, // this is the maximum size that the registry can grow to. If set to `0` (the default) then the registry will not auto grow
		RegistryGrowStep: 32,        // this is how much to step up the registry by each time it runs out of space. The default is `32`.
	})
	defer L.Close()
}
