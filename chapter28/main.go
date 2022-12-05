package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {

	r1 := Rect{
		leftUp:    Point{1, 2},
		rightDown: Point{3, 4},
	}

	// 内存连续分布
	fmt.Printf("r1.leftUp.x 地址=%p\nr1.leftUp.y 地址=%p\nr1.rightDown.x 地址=%p\nr1.rightDown.y 地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// 0xc0000121a0
	// 0xc0000121a8
	// 0xc0000121b0
	// 0xc0000121b8

	// 两*Point类型，本身地址也是连续的
	r2 := Rect2{
		leftUp:    &Point{10, 20},
		rightDown: &Point{30, 40},
	}

	fmt.Printf("r2.leftUp 本身地址=%p\nr2.rightDown 本身地址=%p\n",
		&r2.leftUp, &r2.rightDown)
	// 0xc0000121c0
	// 0xc0000121d0

	// 指向地址不一定连续，系统分配
	fmt.Printf("r2.leftUp 指向地址=%p\nr2.rightDown 指向地址=%p\n",
		r2.leftUp, r2.rightDown)
	// 0xc0000a6080
	// 0xc0000a6090
}
