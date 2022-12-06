package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	//Num float64
	Num int
	//Name string
}

type Student struct {
	Name string
	Age  int
}

type Stu Student

func main() {
	// 两结构体强转，要有相同的字段：字段个数、名称、类型要完全一致
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)

}
