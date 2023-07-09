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

func main() {

	// 反射案例，演示对基本数据类型、interface{}、reflect.Value进行反射的基本操作
	var num int = 100
	reflectTest01(num)
	//reflectType = int
	//reflectValue=100 type=reflect.Value

}
