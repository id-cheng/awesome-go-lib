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
	port := cfg.Section("server").Key("port").MustInt(8888)
	isSSL := cfg.Section("server").Key("is_ssl").MustBool(false)
	fmt.Println("port:", port)
	fmt.Println("is_ssl", isSSL)
}
