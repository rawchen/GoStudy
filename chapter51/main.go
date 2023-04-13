package main

import "fmt"

// 管道的使用
func main() {
	// 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值为=%v, intChan本身地址=%p\n", intChan, &intChan)
	// intChan的值为=0xc00010a080, intChan本身地址=0xc00000a028

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//intChan<- num
	//intChan<- num
	// 管道数据超过容量大小报错
	// fatal error: all goroutines are asleep - deadlock!

	// 输出管道长度和cap容量
	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))
	// channel len = 2, cap = 3

	// 从管道读数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2 = %v\n", num2)
	// num2 = 10
	// 先进先出
	// 输出管道长度和cap容量
	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))
	// channel len = 1, cap = 3
	// 容量不变，长度少一

}
