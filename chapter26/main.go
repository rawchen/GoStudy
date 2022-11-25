package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	Ptr    *int              // 指针
	Slice  []int             // 切片
	Map1   map[string]string // 切片
}

func main() {

	var p1 Person
	fmt.Println(p1)
	// { 0 [0 0 0 0 0] <nil> [] map[]}

	if p1.Ptr == nil {
		fmt.Println("nil Ptr")
	}

	if p1.Slice == nil {
		fmt.Println("nil Slice")
	}

	if p1.Map1 == nil {
		fmt.Println("nil Map1")
	}

	// 使用slice
	p1.Slice = make([]int, 10)
	p1.Slice[0] = 100

	fmt.Println(p1)
}
