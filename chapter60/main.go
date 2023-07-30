package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	reflectType := reflect.TypeOf(b)
	fmt.Println("reflectType =", reflectType)

	reflectValue := reflect.ValueOf(b)
	// 如果需要将反射的值用于计算，它的类型为reflect.Value，需要强制装换为int
	n2 := 2 + reflectValue.Int()
	fmt.Println("n2 =", n2)

	fmt.Printf("reflectValue=%v type=%T\n", reflectValue, reflectValue)

	// 将reflectValue转为interface{}
	iv := reflectValue.Interface()
	//将interface{}断言装成需要的类型
	num2 := iv.(int)
	fmt.Println("num2 =", num2)
}

// 演示反射[结构体]
func reflectTest02(b interface{}) {
	reflectType := reflect.TypeOf(b)
	fmt.Println("reflectType =", reflectType)

	reflectValue := reflect.ValueOf(b)

	// 获取变量对应的Kind
	reflectKind01 := reflectType.Kind()
	reflectKind02 := reflectValue.Kind()
	fmt.Printf("reflectKind01=%v reflectKind02=%v\n", reflectKind01, reflectKind02)
	//reflectKind01=struct reflectKind02=struct

	// 将reflectValue转为interface{}
	iv := reflectValue.Interface()
	//将interface{}断言装成需要的类型
	fmt.Printf("iv=%v, iv type=%T\n", iv, iv)

	// 带检测的类型断言
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {

	// 反射案例，演示对基本数据类型、interface{}、reflect.Value进行反射的基本操作
	var num int = 100
	reflectTest01(num)
	//reflectType = int
	//reflectValue=100 type=reflect.Value

	stu := Student{
		Name: "tom",
		Age:  20,
	}

	reflectTest02(stu)
	//reflectType = main.Student
	//iv={tom 20}, iv type=main.Student

	// reflect包实现了运行时的反射，结构体编译期无法知道实际类型

}
