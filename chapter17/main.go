package main

import "fmt"

func main() {
	num01 := 100
	fmt.Printf("num01的类型=%T, num01的值=%v, num01的地址=%v \n", num01, num01, &num01)

	num02 := new(int) // *int指针类型

	//				*int		0xc0000ac090	0xc0000d8020
	fmt.Printf("num02的类型=%T, num02的值=%v, num02的地址=%v, num02指向的值=%v \n",
		num02, num02, &num02, *num02)

	// 修改num02指向的值为100
	*num02 = 100
	fmt.Println("num02 =", *num02)

}
