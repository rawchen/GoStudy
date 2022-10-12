package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("go on 1")

	err := readConfigFile("config2.ini")
	if err != nil {
		panic(err) // 内置函数
	}
	fmt.Println("go on 2")
}

func test() {
	// 使用defer + recover来捕获和处理异常
	defer func() {
		err := recover() // 内置函数
		if err != nil {
			fmt.Println("err =", err)
		}
	}()
	num01 := 10
	num02 := 0
	fmt.Println(num01 / num02)
}

func readConfigFile(name string) (err error) {
	// 如果传入的不是config.ini就返回自定义错误
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
