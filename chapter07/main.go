package main

import "fmt"

func main() {
	fmt.Println(10 + 7)
	fmt.Println(10 - 7)
	fmt.Println(10 * 7)
	fmt.Println(10 / 7)
	fmt.Println(10 % 7)

	fmt.Println(10 / 3)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	var n3 float32 = 10.0 / 3
	fmt.Println(n3)

	n1++
	fmt.Println(n1)

	//关系运算
	var t1 int = 9
	var t2 int = 8
	fmt.Println(t1 == t2)
	fmt.Println(t1 != t2)
	fmt.Println(t1 > t2)
	fmt.Println(t1 | t2)  //9		//按位或 参与运算的两数各对应的二进制位相或 如果两数对应的二进制位有一个为 1，那么结果为 1， 否则结果为 0。
	fmt.Println(t1 & t2)  //8		//按位与 参与运算的两数各对应的二进制位相与 如果两数对应的二进制位都为 1，那么结果为 1， 否则结果为 0。
	fmt.Println(t1 ^ t2)  //1		//按位异或 参与运算的两数各对应的二进制位相异或 如果两数对应的二进制位不同，那么结果为 1， 否则结果为 0。
	fmt.Println(t1 << t2) //2304	//左移 把运算符 << 左边的运算数的各二进制位全部左移若干位，高位丢弃，低位补 0。 左移 N 位，就是乘以 2 的 N 次方。
	fmt.Println(t1 >> 1)  //4		//右移 把运算符 << 左边的运算数的各二进制位全部右移若干位。 右移 N 位，就是除以 2 的 N 次方。

	fmt.Println(t2 | t1)   //9
	fmt.Println(111 | 222) //255
	fmt.Println(111 & 222) //78

	//逻辑运算
	if t1 > t2 && 2 > 1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if !(t1 < t2) {
		fmt.Println("如果t1不小于t2，输出")
	}

	// 其它运算法
	// & 返回变量存储地址
	// * 指针变量
	a := 100
	fmt.Println("x的地址为：", &a)

	var ptr *int = &a
	fmt.Println("x指向的值为：", *ptr)
}
