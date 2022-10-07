package main

import "fmt"

func main() {
	result := sum(10, 20)
	fmt.Println("main()的result =", result)

}

func sum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer的语句压入到独立的defer栈
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈执行
	// 入栈时，也会将相关值同时拷贝
	defer fmt.Println("n1 =", n1)
	defer fmt.Println("n2 =", n2)
	result := n1 + n2
	n1++
	fmt.Println("sum()的result =", result, "\tn1 =", n1)
	return result
}

// sum()的result = 30      n1 = 11
// n2 = 20
// n1 = 10
// main()的result = 30
