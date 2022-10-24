package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	for i := 0; i < 4; i++ {
		fmt.Println(arr[i])
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}

	var arr02 [5][1]int                         // 因为二维只开辟1个，所以是一个字节
	fmt.Printf("arr02[0]的地址 = %p\n", &arr02[0]) // 0xc00000a540
	fmt.Printf("arr02[1]的地址 = %p\n", &arr02[1]) // 0xc00000a548

	var arr03 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr03)

	for i, v := range arr03 {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, v2)
		}
		fmt.Println()
	}
}
