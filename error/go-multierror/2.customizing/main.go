package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func main() {
	result := new(multierror.Error)
	result.ErrorFormat = func(errors []error) string {
		return "errors!!!"
	}
	fmt.Println(result.Error())

	errorList := []error{
		errors.New("foo"),
		errors.New("bar"),
	}

	multi := &multierror.Error{
		Errors: errorList,
	}
	fmt.Println("error:", multi.Error())
	fmt.Println("errors:", multi.Errors)
	multi.ErrorFormat = func(es []error) string {
		return "foo"
	}
	fmt.Println("error:", multi.Error())
	fmt.Println("orNil:", multi.ErrorOrNil())
}
