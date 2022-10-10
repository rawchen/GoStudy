package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello我"

	fmt.Println(len(str)) // 8

	// 字符串遍历
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符=%c\n", str2[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("12啊")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果：", n)
	}

	// 整数转字符串
	str4 := strconv.Itoa(12345)
	fmt.Println("str4 =", str4)

	// 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes = %v \n", bytes)
	fmt.Printf("bytes = %c \n", bytes)

	// byte转字符串
	str6 := string([]byte{97, 98, 99})
	fmt.Println("str6 =", str6)

	// 10进制转2，8，16进制
	str7 := strconv.FormatInt(123, 2)
	fmt.Println("str7 =", str7)
	str77 := strconv.FormatInt(123, 16)
	fmt.Println("str7 =", str77)

	// 子串包含
	b := strings.Contains("hello world", "he")
	fmt.Println("b =", b)

	// 子串统计
	str9 := strings.Count("abaababa", "aba")
	fmt.Println("str9 =", str9)

	// 不区分大小写的字符串比较
	result10 := strings.EqualFold("abc", "Abc")
	fmt.Println("result10 =", result10)
}
