package main

import (
	"dario.cat/mergo"
	"fmt"
	"log"
)

func main() {
	m1 := make(map[string]interface{})
	m1["nums"] = []uint32{1, 2}

	m2 := make(map[string]interface{})
	m2["nums"] = []int{3}

	if err := mergo.Map(&m1, &m2, mergo.WithOverride, mergo.WithTypeCheck); err != nil {
		log.Println("err:", err)
	}
	if err := mergo.Map(&m1, &m2); err != nil {
		return
	}
	fmt.Println(m1)

	if err := mergo.Map(&m1, &m2, mergo.WithOverride); err != nil {
		return
	}
	fmt.Println(m1)
}
