package main

import "fmt"

func main() {
	// 管道可以只声明为只读或者只写

	// 默认双向的
	// var chan1 chan int

	// 只可写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 3

	fmt.Println("chan2 =", chan2)

	// 无效运算: <-chan2 (从仅发送类型 chan<- int 接收)
	// num := <-chan2

	// 只可读
	var chan3 <-chan int
	num2 := <-chan3

	// 无效运算: chan3 <- 10 (发送到仅接收类型 <-chan int)
	// chan3 <- 10
	fmt.Println("num2 =", num2)
}
