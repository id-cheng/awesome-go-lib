package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string
	Age     int
	Married bool
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func inspectStruct(u any) {
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Uint())

		case reflect.Bool:
			fmt.Printf("field:%d type:%s value:%t\n", i, field.Type().Name(), field.Bool())

		case reflect.String:
			fmt.Printf("field:%d type:%s value:%q\n", i, field.Type().Name(), field.String())

		default:
			fmt.Printf("field:%d unhandled kind:%s\n", i, field.Kind())
		}
	}
}

func inspectSlice(array any) {
	v := reflect.ValueOf(array)
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Println()
}

func inspectMap(m any) {
	v := reflect.ValueOf(m)
	for i, k := range v.MapKeys() {
		field := v.MapIndex(k)
		fmt.Println("i:", i)
		fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
	}
}

func inspectFunc(name string, f any) {
	t := reflect.TypeOf(f)
	fmt.Println(name, "input:")
	for i := 0; i < t.NumIn(); i++ {
		in := t.In(i)
		fmt.Printf("in %d:%s\n", i, in.Name())
	}
	fmt.Println(name, "output:")
	for i := 0; i < t.NumOut(); i++ {
		out := t.Out(i)
		fmt.Printf("out %d:%s\n", i, out.Name())
	}
}

func inspectMethod(o any) {
	t := reflect.TypeOf(o)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		fmt.Println(m)
	}
}

func main() {
	u := User{
		Name:    "cheng",
		Age:     18,
		Married: false,
	}
	inspectStruct(u)

	inspectMap(map[int]int{1: 2, 3: 4})

	inspectSlice([]int{1, 2, 3})

	inspectFunc("Add", func(a, b int) int {
		return a + b
	})
	inspectMethod(&u)
}
