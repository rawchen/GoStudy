package main

import "fmt"

func main() {
	// 数组
	var hens [6]float64
	hens[0] = 1.0
	hens[1] = 5.0
	hens[2] = 2.4
	hens[3] = 40.4
	hens[4] = 3.4
	hens[5] = 3.4

	totalWeight := 0.0
	for i := range hens {
		totalWeight += hens[i]
	}
	// 平均值
	fmt.Printf("%.2f \n", totalWeight/float64(len(hens)))

	// hens的地址为 =0xc00000a420
	// hens[0]地址为=0xc00000a420
	// hens[1]地址为=0xc00000a428
	// hens[2]地址为=0xc00000a430
	fmt.Printf("hens的地址为=%p, hens[0]地址为=%p, hens[1]地址为=%p, hens[2]地址为=%p",
		&hens, &hens[0], &hens[1], &hens[2])

	// 4种初始化方式
	var arr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr01 =", arr01)

	var arr02 = [3]int{4, 5, 6}
	fmt.Println("arr02 =", arr02)

	// 可变容量
	var arr03 = [...]int{7, 8, 9}
	fmt.Println("arr03 =", arr03)

	// 顺序自定义
	var arr04 = [...]int{1: 50, 0: 100, 2: 200}
	fmt.Println("arr04 =", arr04)

	// for-range遍历数组
	names := [...]string{"张三", "李四", "王五"}
	for i, name := range names {
		fmt.Println("索引为 =", i, "姓名为 =", name)
	}

	// 默认情况下是值传递，会进行值拷贝。
	arr05 := [3]int{1, 2, 3}
	test01(arr05)
	fmt.Println(arr05) // [1 2 3]

	arr06 := [3]int{1, 2, 3}
	test02(&arr06)
	fmt.Println(arr06) // [88 2 3]

}

// 值传递测试
func test01(arr [3]int) {
	arr[0] = 88
}

// 引用传递
func test02(arr *[3]int) {
	(*arr)[0] = 88
}
