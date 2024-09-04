package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func main() {
	now := time.Now()
	timestamp := 1600000000
	timeStr := "2008-08-08 20:08:08"

	fmt.Println(cast.ToTime(now))
	fmt.Println(cast.ToTime(timestamp))
	fmt.Println(cast.ToTime(timeStr))

	d, _ := time.ParseDuration("1m30s")
	ns := 30000
	strWithUnit := "130s"
	strWithoutUnit := "130"

	fmt.Println(cast.ToDuration(d))              // 1m30s
	fmt.Println(cast.ToDuration(ns))             // 30µs
	fmt.Println(cast.ToDuration(strWithUnit))    // 2m10s
	fmt.Println(cast.ToDuration(strWithoutUnit)) // 130ns
}
