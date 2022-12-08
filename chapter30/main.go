package main

import "fmt"

type Person struct {
	Name string
}

// test方法只能Person类型变量调用
func (p Person) test() {
	fmt.Println("test方法调用", p.Name)
}

func main() {
	var p Person
	p.Name = "Tom"
	p.test() // 调用结构体方法
}
