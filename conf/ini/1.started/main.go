package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		panic(err)
	}
	mode := cfg.Section("").Key("mode").String()
	data := cfg.Section("paths").Key("data").String()
	protocol := cfg.Section("server").Key("protocol")
	port := cfg.Section("server").Key("port")
	isSSL := cfg.Section("server").Key("is_ssl")
	fmt.Println("mode:", mode)
	fmt.Println("data:", data)
	fmt.Println("protocol:", protocol)
	fmt.Println("port:", port)
	fmt.Println("is_ssl", isSSL)

	cfg.Section("").Key("mode").SetValue("test")
	if err = cfg.SaveTo("my.ini.test"); err != nil {
		panic(err)
	}
}
