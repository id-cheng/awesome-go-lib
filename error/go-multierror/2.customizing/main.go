package main

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
)

func main() {
	result := new(multierror.Error)
	result.ErrorFormat = func(errors []error) string {
		return "errors!!!"
	}
	fmt.Println(result)
}
