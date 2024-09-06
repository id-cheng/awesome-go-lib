package main

import (
	"dario.cat/mergo"
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

var user = User{
	Name: "cheng",
	Age:  18,
}

func main() {
	var u User
	u.Age = 20
	if err := mergo.Merge(&u, user); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user1:", u)

	//传入WithOverride参数，那么目标对象中已经设置的字段也会被覆盖
	if err := mergo.Merge(&u, user, mergo.WithOverride); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user2:", u)

	var m = make(map[string]interface{})
	if err := mergo.Map(&m, user); err != nil {
		log.Fatal(err)
	}
	fmt.Println("map:", m)
}
