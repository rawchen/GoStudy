package main

import "fmt"

//演示golang中指针类型
func main() {
	var i int = 10
	fmt.Println("i的地址=", &i)

	//ptr为指针变量
	//ptr类型为*int
	//ptr本身的值为&i
	var ptr *int = &i
	fmt.Printf("ptr=%v \n", ptr)
	fmt.Printf("ptr的地址=%v \n", &ptr)
	fmt.Printf("ptr指向的值=%v \n", *ptr)
}
