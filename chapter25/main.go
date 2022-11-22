package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
	fmt.Println("函数内：", map1)
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

}
