package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	intFlag      int
	boolFlag     bool
	stringFlag   string
	floatFlag    float64
	durationFlag time.Duration
	flagSet      *flag.FlagSet
)

// init 函数中定义了标志的名称、默认值和描述
func init() {
	flagSet = flag.NewFlagSet("flagSet", flag.ContinueOnError)
	flagSet.IntVar(&intFlag, "int_flag", 0, "int flag value")
	flagSet.Float64Var(&floatFlag, "float_flag", 0, "float flag value")
	flagSet.BoolVar(&boolFlag, "bool_flag", false, "bool flag value")
	flagSet.StringVar(&stringFlag, "string_flag", "", "string flag value")
	flagSet.DurationVar(&durationFlag, "duration_flag", time.Millisecond, "duration flag value")
}

func main() {
	args := []string{
		"-int_flag", "10",
		"-float_flag", "100",
		"-bool_flag", "true",
		"-string_flag", "Hello World",
		"-duration_flag", "1s",
	}

	err := flagSet.Parse(args) // 解析命令行参数
	if err != nil {
		panic(err)
	}
	fmt.Println("int_flag:", intFlag)
	fmt.Println("float_flag:", floatFlag)
	fmt.Println("bool_flag:", boolFlag)
	fmt.Println("string_flag:", stringFlag)
	fmt.Println("duration_flag:", durationFlag)
	fmt.Println("flag_count:", flag.NFlag())
}
