package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go test() // 开启一个协程

	for i := 0; i < 5; i++ {
		fmt.Println("hello02 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
