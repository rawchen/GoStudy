package main

import (
	"fmt"
	"runtime"
)

var (
	myMap = make(map[int]int, 10)
)

// 计算n!
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	myMap[n] = res
}

// 多线程同时占用myMap资源
// fatal error: concurrent map writes
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Printf("cpu数量=%v\n", runtime.NumCPU()) // 12
	runtime.GOMAXPROCS(cpuNum - 4)
	fmt.Printf("cpu数量=%v\n", runtime.NumCPU()) // 12

	for i := 1; i <= 200; i++ {
		go test(i)
	}

	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
