package main

import "fmt"

// 多线程求1-8000内多少素数
func main() {
	intChan := make(chan int, 1000)   // 源
	primeChan := make(chan int, 2000) // 结果
	exitChan := make(chan bool, 2000) // 退出标识

	go putNum(intChan)

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

// 向intChan放入1-8000个数
func putNum(intChan chan int) {
	for i := 0; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}
