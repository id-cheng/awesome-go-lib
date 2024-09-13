package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"reflect"
)

func main() {
	errorList := []error{
		errors.New("foo"),
		errors.New("bar"),
	}
	multi := &multierror.Error{Errors: errorList}
	fmt.Println(reflect.DeepEqual(multi.Errors, multi.WrappedErrors()))
	warpError := errors.Unwrap(multi)
	fmt.Println(warpError)

	multi = nil
	warpError = errors.Unwrap(multi)
	fmt.Println(warpError)             // nil
	fmt.Println(multi.WrappedErrors()) // []
	fmt.Println(multi.Errors)          // panic
}
