package main

import "fmt"

type Person struct {
	Name string
}

// test方法只能Person类型变量调用
func (p Person) test() {
	fmt.Println("test方法调用", p.Name)
}

type Circle struct {
	radius float64
}

// 求面积
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 求面积2
func (c *Circle) area2() float64 {
	c.radius = 10
	return 3.14 * c.radius * c.radius
}

func main() {
	var p Person
	p.Name = "Tom"
	p.test() // 调用结构体方法

	//var c Circle
	//c.radius = 4.0
	//res := c.area()
	//fmt.Println("res =", res)

	var c Circle
	c.radius = 4.0
	res := (&c).area2()
	//res := c.area()
	fmt.Println("res =", res)
	fmt.Println("c.radius =", c.radius)
}
