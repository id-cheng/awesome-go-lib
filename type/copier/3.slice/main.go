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
	users := []User{
		{Name: "jia", Age: 18},
		{Name: "cheng", Age: 20},
	}

	//切片赋值
	employees := make([]Employee, 0, len(users))
	if err := copier.Copy(&employees, &users); err != nil {
		return
	}
	fmt.Println("employees1:", employees)

	//将结构赋值到切片
	employees = make([]Employee, 0, len(users))
	if err := copier.Copy(&employees, &users[0]); err != nil {
		return
	}
	fmt.Println("employees2:", employees)
}
