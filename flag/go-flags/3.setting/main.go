package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

// Option
// required表示对应的选项必须设置值，否则解析时返回ErrRequired错误。
// default用于设置选项的默认值。
type Option struct {
	Required string `short:"r" long:"required" required:"true"`
	Default  string `short:"d" long:"default" default:"default"`
}

func main() {
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	fmt.Println("required: ", opt.Required)
	fmt.Println("default: ", opt.Default)
}

// input go run main.go -r value
// output required:  value
//        default:  default
