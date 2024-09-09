package main

import (
	"fmt"
	"github.com/expr-lang/expr"
)

func main() {
	// 编译
	program, _ := expr.Compile(`2 + 2`)
	// 运行
	output, _ := expr.Run(program, nil)
	fmt.Println(output)
}
