package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	swapMemory, _ := mem.SwapMemory() // 获取交换内存的信息
	fmt.Println(swapMemory.Free)
}
