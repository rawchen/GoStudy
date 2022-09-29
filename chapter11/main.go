package main

import (
	"GoStudy/utils"
	"fmt"
)

func main() {
	utils.PrintOk()

	num := 20
	fmt.Printf("num的地址=%v \n", &num)
	test01(&num)

	result := sum(1, 2, 3, 5)
	fmt.Println(result)
}

func test01(n1 *int) {
	fmt.Printf("n1的地址=%v \n", &n1)
	*n1 = *n1 + 10
	fmt.Println("n1 = ", *n1)

}

func sum(n1 int, args ...int) int {
	sum := n1
	// 遍历args
	for i := range args {
		sum += args[i]
	}
	return sum
}
