package main

import "fmt"

func main() {
	// 切片的基本使用 方式1
	var arr [5]int = [...]int{1, 22, 33, 45, 15}
	// 声明定义一个切片
	slice := arr[1:3]
	// 从数组下标1开始，3结束不包含 -> index[1, 3)
	fmt.Println("slice为 =", slice)              // [22 33]
	fmt.Printf("slice类型为 = %T \n", slice)       // []int
	fmt.Printf("slice元素个数 = %v \n", len(slice)) // 2
	fmt.Printf("slice元素容量 = %v \n", cap(slice)) // 4

	// 切片的基本使用 方式2
	// 演示切片的使用make
	var slice02 []float64 = make([]float64, 5, 10)
	slice02[1] = 10.1
	slice02[3] = 20.8
	// 对于切片，必须make使用
	fmt.Println("slice02为 =", slice02)
	fmt.Printf("slice02元素个数 = %v \n", len(slice02)) // 5
	fmt.Printf("slice02元素容量 = %v \n", cap(slice02)) // 10

	// 切片的基本使用 方式3
	var slice03 []string = []string{"tom", "jack", "mary"}
	fmt.Println("slice03为 =", slice03)
	fmt.Printf("slice03元素个数 = %v \n", len(slice03)) // 3
	fmt.Printf("slice03元素容量 = %v \n", cap(slice03)) // 3

	// 常规for循环遍历
	var arr03 [5]int = [...]int{10, 20, 30, 40, 50}
	slice04 := arr03[1:4]
	for i := 0; i < len(slice04); i++ {
		fmt.Printf("slice04[%v]=%v ", i, slice04[i])
	}
	fmt.Println()

	// for-range
	for i, v := range slice04 {
		fmt.Printf("i=%v v=%v\n", i, v)
	}
	// 演示切片简写 [20, 30, 40]
	slice05 := slice04[:]
	slice06 := slice04[1:]
	slice07 := slice04[:2]
	fmt.Println(slice05)
	fmt.Println(slice06)
	fmt.Println(slice07)

}
