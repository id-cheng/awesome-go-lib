package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug message"`
}

/*
short和long设置短、长选项名字
description设置帮助信息。
命令行传参时，短选项前加-，长选项前加--
*/
func main() {
	var opt Option
	flags.Parse(&opt)

	fmt.Println(opt.Verbose)
}

// input  go run main.go -v -v
// output   [true true]

// input  go run main.go --verbose --verbose
// output   [true true]

// input  go run main.go -v --verbose -v
// output   [true true true]

// input  go run main.go -vvv
// output   [true true true]
