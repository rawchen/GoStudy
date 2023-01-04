package main

import "fmt"

type AInterface interface {
	Say()
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("Student Say()")
}

// 接口本身不能创建实例，但可以指向一个实现了该接口的自定义类型的变量（实例）
func main() {
	// var a AInterface
	// a.Say()
	// invalid memory address or nil pointer dereference

	var stu Student
	var b AInterface = stu
	b.Say()
}
