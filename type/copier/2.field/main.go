package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name  string
	Age   int
	Email string
}

// Phone Copy会调用这个方法，将返回值赋值给目标对象中的字段
func (u *User) Phone() string {
	return "123456"
}

type Employee struct {
	Name  string
	Phone string
	Role  string
}

// Email Copy会以源对象的这个字段作为参数调用目标对象的该方法
func (u *Employee) Email(email string) {
	fmt.Println("email:", email)
}

func main() {
	user := User{Name: "cheng", Age: 18, Email: "123456@gmail.com"}
	employee := new(Employee)

	if err := copier.Copy(employee, &user); err != nil {
		return
	}
	fmt.Println("phone:", employee.Phone)
}
