package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	intFlag      *int
	boolFlag     *bool
	stringFlag   *string
	floatFlag    *float64
	durationFlag *time.Duration
)

// init 函数中定义了标志的名称、默认值和描述
func init() {
	intFlag = flag.Int("int_flag", 0, "int flag value")
	floatFlag = flag.Float64("float_flag", 0, "float flag value")
	boolFlag = flag.Bool("bool_flag", false, "bool flag value")
	stringFlag = flag.String("string_flag", "", "string flag value")
	durationFlag = flag.Duration("duration_flag", time.Millisecond, "duration flag value")
}

/*
go run main.go -int_flag=10 -float_flag=100 -bool_flag=true -string_flag="Hello World" -duration_flag=1s arg1 arg2

	int_flag: 10
	float_flag: 10
	bool_flag: true
	string_flag: Hello World
	duration_flag: 1s
	flag_count: 5
	flag_args: [arg1 arg2]
*/
func main() {
	flag.Parse() // 解析命令行
	fmt.Println("int_flag:", *intFlag)
	fmt.Println("float_flag:", *floatFlag)
	fmt.Println("bool_flag:", *boolFlag)
	fmt.Println("string_flag:", *stringFlag)
	fmt.Println("duration_flag:", *durationFlag)
	fmt.Println("flag_count:", flag.NFlag())
	fmt.Println("flag_args:", flag.Args())
	time.Sleep(*durationFlag)
}
