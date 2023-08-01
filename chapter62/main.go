package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改num值
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	// Elem 返回接口 v 包含的值或指针 v 指向的值。
	// 如果 v 的Kind不是Interface或Pointer ，它会发生恐慌。如果 v 为零，则返回零值。
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)
	//rVal kind=ptr
	//num= 20

	// 理解rVal.Elem()
	//num := 9
	//ptr *int = &num
	//num2 := *ptr
	//fmt.Println("num2=", num2)
}
