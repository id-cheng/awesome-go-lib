package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Note
	Name    string
	Age     int `ini:"age"`
	Male    bool
	Born    time.Time
	Created time.Time `ini:"-"`
}

func main() {
	cfg, _ := ini.Load("my.ini")
	p1 := new(Person)
	err := cfg.MapTo(p1)
	if err != nil {
		return
	}
	fmt.Println("person1", p1)

	var p2 Person
	if err = ini.MapTo(&p2, "my.ini"); err != nil {
		return
	}
	fmt.Println("person2", p2)

	node := &Note{}
	if err = cfg.Section("Note").MapTo(node); err != nil {
		return
	}
	fmt.Println("note", node)

	p2.Age = 35
	if err = ini.ReflectFrom(cfg, &p2); err != nil {
		return
	}
	if err = cfg.SaveTo("person.ini"); err != nil {
		return
	}
}
