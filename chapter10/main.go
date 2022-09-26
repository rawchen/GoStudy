package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	for i := 0; i < 10; i++ {
		// 0-9
		fmt.Println(i)
	}

	fmt.Println("==============")
	for i := 1; i <= 10; i++ {
		// 1-10
		fmt.Println(i)
	}

	fmt.Println("//////////////")
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("!!!!!!!!!!!!!!!")
	// go语言无do-while，使用for{break}实现
	k := 1
	for {
		if k <= 10 {
			fmt.Println(k)
		} else {
			break
		}
		k++
	}

	fmt.Println("ggggggggggggggggg")
	// 字符串遍历
	// 传统遍历中文乱码，会按照字节遍历！
	var txt string = "hello world啊"
	for i := 0; i < len(txt); i++ {
		//fmt.Println(txt[i])
		fmt.Printf("%c \n", txt[i])
	}

	fmt.Println("ooooooooooooooooo")
	// 字符串遍历
	// for-range按照字符遍历，所以不会乱码。
	var txt2 string = "hello world啊"
	for index, val := range txt2 {
		fmt.Printf("index = %d, val = %c \n", index, val)
	}

	fmt.Println("zzzzzzzzzzzzzzzzz")
	// 字符串遍历
	// 切片解决传统遍历中文乱码
	var txt3 string = "hello world哈哈哈哈呵呵呵"
	txt4 := []rune(txt3)
	for i := 0; i < len(txt); i++ {
		fmt.Printf("%c \n", txt4[i])
	}

	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("count = ", count)

	// goto
	fmt.Println("OK1")
	goto label1
	fmt.Println("OK2")
	fmt.Println("OK3")
label1:
	fmt.Println("OK4")
}
