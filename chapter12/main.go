package main

import "fmt"

var (
	// Fun1 全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 匿名函数测试
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 11)

	fmt.Println(result)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	result02 := a(10, 30)
	fmt.Println(result02)

	result03 := Fun1(2, 3)
	fmt.Println(result03)
}
