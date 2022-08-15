package main

import "fmt"

//定义全局变量
var y1 = 100
var y2 = 200
var name3 = "tom"

var (
	y3 = 300
	y4 = 400
)

func getSumAndSub02(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	num01, _ := getSumAndSub02(1, 2)
	fmt.Println(num01)

	var a int
	fmt.Println("a设默认值：", a)

	var b = 0
	fmt.Println("b类型推导：", b)

	c := 0
	fmt.Println("c定义赋值：", c)

	//多变量一次性声明
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//一次性声明多种类型变量
	var t1, name, t3 = 100, "tom", true
	fmt.Println(t1, name, t3)

	//类型推导
	t11, name1, t33 := 100, "tom", true
	fmt.Println(t11, name1, t33)

	fmt.Println(y1, y2, name3)

	fmt.Println(y3, y4)
}
