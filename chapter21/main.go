package main

import "fmt"

// 值传递 引用传递
func main() {
	arr := [5]int{23, 45, 12, 1, 60}
	BubbleSort(&arr)
	fmt.Println("排序后arr =", arr)
}

func BubbleSort(arr *[5]int) {
	arr[0] = 90
	arr[1] = 10
	fmt.Println("排序前arr =", *arr)
}
