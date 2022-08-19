package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int = 100
	n = 10000
	//n = 1.12
	//n := 100

	fmt.Println(n)

	x := 1
	y := 2
	fmt.Println(x + y)
	//fmt.Println("1" + 2)
	fmt.Println("1" + "2")

	//基本数据类型
	//数值-整数 int int8 int16 int32(rune) int64 uint uint8(byte) uint16 uint32 uint64
	//数值-浮点 float32 float64
	//布尔 bool
	//字符串string

	//复杂数据类型
	//指针 Pointer
	//数组
	//结构体 struct
	//管道 Channel
	//函数
	//切片 slice
	//接口 interface
	//map

	//溢出
	//var j int8 = -129
	//fmt.Println("j=", j)

	var l byte = 255
	fmt.Println("l=", l)

	//查看数据类型、整形占用
	var k int16 = -129
	fmt.Println("k=", k)
	fmt.Printf("k数据类型：%T，占用字节：%d   \n", k, unsafe.Sizeof(k))

	var price float32 = 89.12
	fmt.Println("price=", price)

	//精度问题 float64精度更高
	var m, o = 73.1, 13.1
	fmt.Println(m, o)
	fmt.Println(m - o)

	var p = .56
	fmt.Println(p)

	var q = 5.123e2
	fmt.Println(q)

	//字符类型
	//var r byte = "aaa"
	//var r byte = 'aaa'
	var r byte = 'a'
	var s byte = 97
	var t byte = '\t'
	var u int = '你'
	//输出码值:97 97 9 20320
	fmt.Println(r, s, t, u)
	//格式化:a a      你
	fmt.Printf("%c %c %c %c", r, s, t, u)

	//溢出 constant 27979 overflows byte !
	//var v byte = '测'
	//fmt.Printf("%c", v)
}
