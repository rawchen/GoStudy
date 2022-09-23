package main

import "fmt"

func main() {
	var name int = 10

	if name > 1 && name <= 10 {
		fmt.Println(name, "!")
	} else if name > 10 && name <= 20 {
		fmt.Println(name, "!!")
	} else {
		fmt.Println(name)
	}

	var status = 4
	statusT := 1
	switch status { // 常量 变量 函数
	// 常量 变量 函数
	case 1:
		fmt.Println(1)
	case 4, 5:
		fmt.Println("4 or 5")
	case statusT + 1:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	// switch后也可不带表达式 类似if-else
	// 也可范围判断
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age = 10")
	case age == 20:
		fmt.Println("age = 20")
	case age >= 30 && age <= 50:
		fmt.Println("age >= 30 && age <= 50")
	default:
		fmt.Println("no")
	}

	// switch穿透
	// 结果为:10
	//       20
	var num int = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("no")
	}
}
