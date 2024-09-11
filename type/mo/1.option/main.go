package main

import (
	"fmt"
	"github.com/samber/mo"
)

func main() {
	option1 := mo.Some(10)
	option2 := mo.None[int]()
	fmt.Println(option1.Get())
	fmt.Println(option2.Get())
	fmt.Println(option1.IsPresent())
	fmt.Println(option2.OrElse(100))

	option3 := option1.FlatMap(func(val int) mo.Option[int] {
		return mo.Some(val * 2)
	})
	fmt.Println(option3.Get())

	option4 := option1.Match(func(val int) (int, bool) {
		return val * 10, true
	}, func() (int, bool) {
		return 0, false
	})
	fmt.Println(option4.Get())
}
