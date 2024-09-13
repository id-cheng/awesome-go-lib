package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func main() {
	err := new(multierror.Error)
	fmt.Println(err.ErrorOrNil())
	err.Errors = []error{errors.New("foo")}
	fmt.Println(err.ErrorOrNil())
}
