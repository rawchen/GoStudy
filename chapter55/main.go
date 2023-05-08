package main

import (
	"fmt"
	"time"
)

// 多线程求1-x内多少素数
// 8000 	44.1836ms
// 50000 	209.0029ms
func main() {
	var x = 50000
	bT := time.Now()

	intChan := make(chan int, x)       // 源
	primeChan := make(chan int, 10000) // 结果
	exitChan := make(chan bool, 4)     // 退出标识

	go putNum(intChan, x)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// 当我们从exitChan取出四个结果就可以放心关闭primeChan
		close(primeChan)
	}()

	var count = 0
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		count++
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Printf("素数个数为=%d\n", count)
	fmt.Println("main线程退出")

	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			// 管道取不出东西了
			break
		}
		flag = true // 假设为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			// 如果都循环一遍发现都除不尽就是素数
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据退出")
	// 还不能关闭primeChan，向exitChan写入true
	exitChan <- true
}

// 向intChan放数据
func putNum(intChan chan int, x int) {
	for i := 0; i <= x; i++ {
		intChan <- i
	}
	close(intChan)
}
