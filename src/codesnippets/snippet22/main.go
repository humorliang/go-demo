package main

import (
	"runtime"
	"fmt"
	"time"
)

/*
Print The Current Memory Usage
输出当前的内存使用率
*/
func main() {
	OutMemUage()
	var overall [][]int

	for i := 0; i < 5; i++ {
		a := make([]int, 0, 99999)
		overall = append(overall, a)
		OutMemUage()
		time.Sleep(time.Second)
	}
	//清理内存对象
	overall = nil
	OutMemUage()

	//清理GC
	runtime.GC()
	OutMemUage()
}

func OutMemUage() {
	//定义内存状态变量
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	//输出内存信息
	//Alloc是已分配堆对象的字节
	fmt.Printf("Alloc = %v MiB", BitToMiB(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", BitToMiB(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", BitToMiB(m.Sys))
	fmt.Printf("\tNumGC = %v MiB \n", m.NumGC)
}

//bit 转 Mib
func BitToMiB(b uint64) uint64 {
	return b / 1024 / 1024
}
