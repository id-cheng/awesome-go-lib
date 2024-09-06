package main

import (
	"dario.cat/mergo"
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
	Nums []int
}

var user = User{
	Name: "cheng",
	Age:  18,
}

func main() {
	var u User
	u.Nums = []int{1, 2, 3}
	if err := mergo.Merge(&u, user, mergo.WithOverride); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user1:", u)

	user.Nums = []int{4, 5, 6}
	if err := mergo.Merge(&u, user, mergo.WithOverride, mergo.WithAppendSlice); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user2:", u)

	if err := mergo.Merge(&u, user, mergo.WithOverride, mergo.WithOverrideEmptySlice); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user3:", u)
}
