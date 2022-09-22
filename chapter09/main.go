package main

import "fmt"

func main() {
	var i int = 20
	// 二进制
	fmt.Printf("%b\n", i)

	// 0开头则为8进制，输出为10进制的9
	var j int = 011
	fmt.Println("j=", j)

	// 0x或0X开头则为16进制，输出为10进制的229
	var k int = 0xe5
	fmt.Println("k=", k)
}
