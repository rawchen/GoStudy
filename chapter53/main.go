package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// 此时无法再写入数据到channel
	//intChan <- 300

	n1 := <-intChan
	fmt.Println("n1 =", n1)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// 遍历管道不能使用普通for
	//for i := 0; i < len(intChan2); i++ {
	//
	//}

	close(intChan2)
	// 没有关闭就遍历会报错
	// fatal error: all goroutines are asleep - deadlock!
	for v := range intChan2 {
		fmt.Printf("v = %v\n", v)
	}
}
