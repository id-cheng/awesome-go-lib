package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

// Option
// 所有的基本类型和它们的切片
// map只支持键为string,值为基础类型
// 函数类型
type Option struct {
	IntFlag   int               `short:"i" long:"int" description:"int flag value"`
	BoolFlag  bool              `long:"bool" description:"bool flag value"`
	FloatFlag float64           `long:"float" description:"float64 flag value"`
	StrFlag   string            `short:"s" long:"str" description:"string flag value"`
	StrSlice  []string          `long:"strSlice" description:"string slice flag value"`
	PStrSlice []*string         `long:"pstrSlice" description:"slice of pointer of string flag value"`
	Call      func(string)      `long:"call" description:"callback"`
	Map       map[string]string `long:"map" description:"A map from string to string"`
}

func main() {
	var opt Option
	opt.Call = func(value string) {
		fmt.Println("in callback: ", value)
	}

	if _, err := flags.Parse(&opt); err != nil {
		log.Println("Parse error:", err)
		return
	}

	fmt.Println("int flag:", opt.IntFlag)
	fmt.Println("bool flag:", opt.BoolFlag)
	fmt.Println("float flag:", opt.FloatFlag)
	fmt.Println("str flag:", opt.StrFlag)
	fmt.Println("str slice:", opt.StrSlice)
	fmt.Println("map flag:", opt.Map)
	for i, v := range opt.PStrSlice {
		fmt.Printf("%d:%v\n", i, *v)
	}
}

// input  go run main.go --map key1:1 --map key2:2
// output map flag: map[key1:1 key2:2]

// input  go run main.go --call test1
// output in callback:  test1

// input go run  --pstrSlice s1 --pstrSlice s2
// output 0:s1\n 1:s2
