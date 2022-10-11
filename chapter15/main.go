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

	// 子串第一次出现的位置，没有返回-1
	int11 := strings.Index("abcdef", "bcd")
	fmt.Println("int11 =", int11)

	// 子串最后一次出现的位置，没有返回-1
	int12 := strings.LastIndex("abcdefabc", "abc")
	fmt.Println("int12 =", int12) // 6

	// 子串替换，n则为替换几个，-1全部替换
	str13 := strings.Replace("abcdefabc", "abc", "A", -1)
	fmt.Println("str13 =", str13) // AdefA

	// 按某字符分割
	arr14 := strings.Split("123,456,789", ",")
	fmt.Println("arr14 =", arr14)

	// 大小写转换
	str15 := strings.ToLower("Abc") // ToUpper()
	fmt.Println("str15 =", str15)

	// 去两边空格
	str16 := strings.TrimSpace(" abc ")
	fmt.Println("str16 =", str16)

	// 去两边指定字符
	str17 := strings.Trim("1abc21", "1") // TrimLeft() TrimRight()
	fmt.Println("str17 =", str17)

	// 以指定字符串开头
	result20 := strings.HasPrefix("123456", "123")
	fmt.Println("result20 =", result20)

	// 以指定字符串结尾
	result21 := strings.HasSuffix("123456", "456")
	fmt.Println("result21 =", result21)
}
