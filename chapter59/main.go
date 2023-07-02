package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	// 这里我们可以使用defer + recover
	defer func() {
		// 捕获test抛出的panic异常，但不结束程序
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	// 定义一个map
	var myMap map[int]string
	myMap[0] = "golang" // 没有make 将error
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() =", i)
		time.Sleep(time.Second)
	}
	// panic: assignment to entry in nil map

}
