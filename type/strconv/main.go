package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("10")
	s := strconv.Itoa(-10)
	fmt.Println("int:", i)
	fmt.Println("str:", s)

	i64, _ := strconv.ParseInt("-10", 10, 64)
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	u, _ := strconv.ParseUint("10", 10, 64)
	fmt.Println("int64", i64)
	fmt.Println("bool", b)
	fmt.Println("float64", f)
	fmt.Println("uint", u)

	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := fmt.Sprintf("%v", 3.1415)
	s4 := strconv.FormatInt(-42, 16)
	s5 := strconv.FormatUint(42, 16)
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)
	fmt.Println("s4", s5)

	fmt.Println(strconv.Quote("hello 世界")) // 带引号
	fmt.Println(strconv.QuoteToASCII("hello 世界"))
}
