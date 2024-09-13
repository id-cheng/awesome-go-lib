package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
)

// 构建错误列表
func main() {
	err := errors.New("err")
	result := multierror.Append(err, errors.ErrUnsupported)
	fmt.Println(result.Errors)
	fmt.Println(result.Error())
}
