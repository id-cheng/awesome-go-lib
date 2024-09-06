package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	user := User{Name: "cheng", Age: 18}
	employee := Employee{}
	if err := copier.Copy(&employee, &user); err != nil {
		return
	}
	fmt.Println(employee)
}
