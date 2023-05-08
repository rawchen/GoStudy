package main

import (
	"fmt"
	"time"
)

// 一线程求1-x内多少素数
// 8000 	48.6127ms
// 50000 	500.7248ms
func main() {
	bT := time.Now()

	var x = 50000
	var count int
	var flag bool
	for i := 1; i <= x; i++ {
		flag = true // 假设为素数
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}

		if flag {
			count++
			fmt.Printf("素数=%d\n", i)
		}
	}

	fmt.Printf("素数个数为=%d\n", count)
	fmt.Println("main线程退出")

	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
}
