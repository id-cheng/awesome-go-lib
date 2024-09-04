package main

import (
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		panic(err)
	}
	cfg.Section("").Key("mode").SetValue("test")
	if err = cfg.SaveTo("my.ini.test"); err != nil {
		panic(err)
	}
}
