package main

import "fmt"

func main() {
	// 切片的基本使用
	var arr [5]int = [...]int{1, 22, 33, 45, 15}
	// 声明定义一个切片
	slice := arr[1:3]
	// 从数组下标1开始，3结束不包含 -> index[1, 3)
	fmt.Println("slice为 =", slice)              // [22 33]
	fmt.Printf("slice类型为 = %T \n", slice)       // []int
	fmt.Printf("slice元素个数 = %v \n", len(slice)) // 2
	fmt.Printf("slice元素容量 = %v \n", cap(slice)) // 4
}
