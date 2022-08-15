package main

import "fmt"

func testFunction(num1 int, num2 int) int {
	return (num1 + num2)
}

func testPtr(num *int) {
	*num = 20
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

//java基础差别
//区分大小写
//定义必须使用
//常用行注释
//一行一语句，无需分号
func main() {
	fmt.Println(123)
	var a = 10
	fmt.Println("输出：%v", a)
	fmt.Println("输出：%v", &a)
	//fmt.Printf("test: %v", testPtr())

	fmt.Println("%v", testFunction(2, 2))

	r1, r2 := getSumAndSub(1, 2)

	fmt.Println("%v", r1)
	fmt.Println("%v", r2)

	fmt.Println("123232313123123333333333333333313",
		"1233333333333333333333333333333333333",
		"22222222222222222222222222222222222")

}
