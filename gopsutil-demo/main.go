package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"time"
)

func main() {
	v, _ := mem.VirtualMemory()

	per, _ := cpu.Percent(time.Second, false)

	pro, _ := process.NewProcess(9100)
	cpuPercent, _ := pro.CPUPercent()
	fmt.Println(cpuPercent)

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	fmt.Println(per)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
