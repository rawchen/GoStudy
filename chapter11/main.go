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
}

func test01(n1 *int) {
	fmt.Printf("n1的地址=%v \n", &n1)
	*n1 = *n1 + 10
	fmt.Println("n1 = ", *n1)

}
