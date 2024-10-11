package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"time"
)

func main() {
	physicalCnt, _ := cpu.Counts(false) // 物理核心
	logicalCnt, _ := cpu.Counts(true)   // 逻辑核心
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	perPercents, _ := cpu.Percent(3*time.Second, true)
	fmt.Printf("total percent:%v per percents:%v", totalPercent, perPercents)

	info, _ := cpu.Info()
	fmt.Println(info)

	times, _ := cpu.Times(true) //false返回总的  true返回单个的
	fmt.Println(times)
}
