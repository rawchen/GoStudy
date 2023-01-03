package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}

type Test struct {
	Monster
	int
	Age int
}

type Parent struct {
	Monster
	Test
}

func main() {
	var test Test
	test.Name = "名称"
	test.Age = 20
	test.int = 30
	fmt.Println("test=", test)

	var p Parent
	// 不明确的引用
	//fmt.Println("p.Age=", p.Age)
	fmt.Println("p.Age=", p.Test.Age)
}
