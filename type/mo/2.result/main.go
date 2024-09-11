package main

import (
	"errors"
	"fmt"
	"github.com/samber/mo"
)

func main() {
	ok := mo.Ok(10)
	result1 := ok.OrElse(100)
	err1 := ok.Error()
	fmt.Println(result1, err1)

	ko := mo.Err[int](errors.ErrUnsupported)
	result2 := ko.OrElse(100)
	err2 := ko.Error()
	fmt.Println(result2, err2)
}
