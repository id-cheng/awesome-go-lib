package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
	"time"
)

func main() {
	timestamp, _ := host.BootTime() // 开机时间
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))

	version, _ := host.KernelVersion()
	fmt.Println(version)
	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)

	users, _ := host.Users()
	for _, user := range users {
		fmt.Println(user.String())
	}
}
