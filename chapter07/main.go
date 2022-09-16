package main

import "fmt"

func main() {
	fmt.Println(10 + 7)
	fmt.Println(10 - 7)
	fmt.Println(10 * 7)
	fmt.Println(10 / 7)
	fmt.Println(10 % 7)

	fmt.Println(10 / 3)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	var n3 float32 = 10.0 / 3
	fmt.Println(n3)

	n1++
	fmt.Println(n1)

	//关系运算
	var t1 int = 9
	var t2 int = 8
	fmt.Println(t1 == t2)
	fmt.Println(t1 != t2)
	fmt.Println(t1 > t2)
	fmt.Println(t1 | t2)
	fmt.Println(t1 & t2)
	fmt.Println(t1 << t2)
	fmt.Println(t1 >> 1)
}
