package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
	fmt.Println("函数内：", map1)
}

type Student struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// map为引用类型
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)

	fmt.Println("主函数：", map1)

	// 结构体
	students := make(map[string]Student)
	//student1 := Student{
	//	Name:    "张三",
	//	Age:     22,
	//	Address: "深圳",
	//}
	student1 := Student{"张三", 22, "深圳"}
	student2 := Student{"李四", 25, "厦门"}
	students["s1"] = student1
	students["s2"] = student2
	fmt.Println(students)

	// 遍历学生map
	for k, v := range students {
		fmt.Printf("学生编号是：%v \n", k)
		fmt.Printf("结构体内容：%v \n", v)
		fmt.Printf("学生姓名是：%v \n", v.Name)
		fmt.Printf("学生年龄是：%v \n", v.Age)
		fmt.Printf("学生地址是：%v \n", v.Address)
	}

	fmt.Printf("student1的地址=%p\n", &student1) // 0xc00007e570

}
