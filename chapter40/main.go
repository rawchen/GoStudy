package main

import "fmt"

type Point struct {
	x int
	y int
}

// 将a赋给Point变量
func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	// 不可以
	// b = a

	var b Point
	b = a.(Point)

	fmt.Println(b)
}
