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

	var b Point

	// 不可以
	// b = a

	b = a.(Point)
	fmt.Println(b)

	// 类型断言其它案例
	//var x interface{}
	//var b2 float32 = 1.1
	//x = b2 // 空接口可接任意类型
	//y := x.(float32)
	//fmt.Printf("y 的类型是 %T 值是=%v", y, y)

	// 类型断言(带检测)
	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口可接任意类型
	y, ok := x.(float64)

	if ok == true {
		fmt.Println("covert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("covert fail")
	}

	fmt.Println("go on...")
}
