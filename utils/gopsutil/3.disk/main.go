package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
)

func main() {
	diskIO, _ := disk.IOCounters()
	for name, stat := range diskIO {
		fmt.Println(name)
		fmt.Println(stat)
	}

	partitions, _ := disk.Partitions(false) //分区
	for _, info := range partitions {
		fmt.Println(info)
	}

	usage, _ := disk.Usage("C:/Users")
	fmt.Println(usage)
}
