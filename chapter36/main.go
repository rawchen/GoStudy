package main

import "fmt"

type Student struct {
	Name string
}

type AInterface interface {
	// 接口中不能有任何变量
	//Name string
	Say()
}

func (stu Student) Say() {
	fmt.Println("Student Say()")
}

type BInterface interface {
	Speak()
}

func (stu Student) Speak() {
	fmt.Println("Teacher Speak()")
}

// 同时实现多个接口
func main() {
	var stu Student
	var a AInterface = stu
	var b BInterface = stu
	a.Say()
	b.Speak()
}
