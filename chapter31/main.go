package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", s.Name, s.Age)
	return str
}

// 真正决定是值拷贝还是地址拷贝，看方法和哪个类型绑定，比如地址拷贝 s *Student
func main() {
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(&stu)
}
