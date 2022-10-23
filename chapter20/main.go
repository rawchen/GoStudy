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

	// 因为切片是引用底层数组，所以一改都改
	slice07[0] = 123
	fmt.Println(slice07) // [123, 30]
	fmt.Println(slice04) // [123, 30, 40]

	// 动态追加
	var slice08 []int = []int{1, 2, 3}
	slice09 := append(slice08, 4)
	fmt.Println("slice08 =", slice08) // [1 2 3]
	fmt.Println("slice09 =", slice09) // [1 2 3 4]

	slice10 := append(slice09, slice09...)
	fmt.Println("slice10 =", slice10) // [1 2 3 4 1 2 3 4]

	// 拷贝
	var slice11 = []int{1, 2, 3}
	slice12 := make([]int, 5)
	fmt.Println("slice11 =", slice11) // [1 2 3]
	fmt.Println("slice12 =", slice12) // [0 0 0 0 0]
	copy(slice12, slice11)
	fmt.Println("slice11 =", slice11) // [1 2 3]
	fmt.Println("slice12 =", slice12) // [1 2 3 0 0]

	// string切片
	slice13 := "123456"
	slice14 := slice13[2:]
	fmt.Println("slice14 =", slice14) // [3 4 5 6]

	slice15 := []byte(slice13) // 如果处理中文则为[]rune
	slice15[0] = '9'
	slice16 := string(slice15)
	fmt.Println("slice16 =", slice16) // [9 2 3 4 5 6]
}
