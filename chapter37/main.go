package main

import "fmt"

type AInterface interface {
	// 接口可以继承接口
	BInterface
	CInterface
	test01()
}

type BInterface interface {
	test02()
}

type CInterface interface {
	test03()
}

type Student struct {
}

func (stu Student) test01() {
	fmt.Println("test01()")
}

func (stu Student) test02() {
	fmt.Println("test02()")
}

// 将test03()注释后提示：
// 无法将 'stu' (类型 Student) 用作类型 AInterface 类型未实现 'AInterface'，
// 因为缺少某些方法: test03()

func (stu Student) test03() {
	fmt.Println("test03()")
}

// A接口继承了B、C接口，如果要实现A接口，就需要将B、C接口的方法都实现
func main() {
	var stu Student
	var a AInterface = stu
	a.test01()
}
