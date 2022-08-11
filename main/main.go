package main

import "fmt"

func testFunction(num1 int, num2 int) int {
	return (num1 + num2)
}

func testPtr(num *int) {
	*num = 20
}

func main() {
	fmt.Println(123)
	var a = 10
	fmt.Printf("输出：%v", a)
	fmt.Printf("输出：%v", &a)
	//fmt.Printf("test: %v", testPtr())

	fmt.Printf(string(testFunction(1, 2)))
}
