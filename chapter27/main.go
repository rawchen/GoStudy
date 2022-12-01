package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p2 := Person{"张三", 21}
	fmt.Println(p2)

	var p3 *Person = new(Person)
	// p3为指针，标准赋值方法
	//(*p3).Name = "123"
	p3.Name = "123"
	(*p3).Age = 88

	fmt.Println(*p3)

	var p4 *Person = &Person{}
	//var p4 *Person = &Person{"tom", 30}
	// (*p4).Name = "tom"
	p4.Name = "tom"
	fmt.Println(p4)  // &{tom 0}
	fmt.Println(&p4) // 0xc0000ca020
	fmt.Println(*p4) // {tom 0}

	// 结构体内存分配
	var p5 Person
	p5.Name = "李四"
	p5.Age = 70

	var p6 *Person = &p5

	fmt.Println(p5.Name)
	fmt.Println(p6.Name)

	fmt.Printf("p5的地址为%p\n", &p5)             // 0xc0000080f0
	fmt.Printf("p6的地址为%p p6的值为%p\n", &p6, p6) // 0xc00000a038  0xc0000080f0

}
