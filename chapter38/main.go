package main

import "fmt"

type T interface {
}

type Student struct {
}

// 可以把任何一个变量赋给空接口（所有类型都实现了空接口）
func main() {
	var stu Student
	var t T = stu
	fmt.Println(t)

	var t2 interface{} = stu
	fmt.Println(t2)

	var num1 float64 = 8.8
	t = num1
	fmt.Println(t2, t)
}
