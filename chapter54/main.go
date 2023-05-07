package main

import (
	"fmt"
	"time"
)

// goroutine和channel协同工作
// 开启一个writeData协程，向管道intChan写入20个整数
// 开启一个readData协程，从管道intChan中读取writeData写入的数据
// 主线程需要等待writeData和readData协程都完成工作才能退出

func writeData(intChan chan int) {
	for i := 1; i <= 20; i++ {
		intChan <- i
		fmt.Println("writeData ", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)
	// 如果注掉以上读数据线程，则会在intChan发生阻塞，出现dead lock

	//time.Sleep(time.Second * 10)

	//for {
	//	_, ok := <- exitChan
	//	if !ok {
	//		break
	//	}
	//}

	// 管道阻塞主进程，直到管道能取到数据
	_ = <-exitChan
}
