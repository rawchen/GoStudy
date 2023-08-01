package main

import "fmt"

func main() {
	var num int
	num = 9
	// 常量声明时必须赋值
	const tax int = 0
	// 常量是不能修改的
	//tax = 10

	// 常量只能修饰bool/数值型（int、float系列）/string类型
	fmt.Println(num, tax)

	//const b = num / 3
	const (
		a = iota
		b
		c
		d
	)
	// 0 1 2 3
	fmt.Println(a, b, c, d)

	const (
		e    = iota
		f    = iota
		g, h = iota, iota
	)
	// 0 1 2 2
	fmt.Println(e, f, g, h)

	// 仍然通过首字母大小写控制常量访问范围
}
