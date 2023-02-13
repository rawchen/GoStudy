package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 一次性往内存读取文件
func main() {
	file := "d:/test.txt"

	// 弃用
	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err = %v", err)
	}
	fmt.Printf("%v\n", string(readFile))

	// 新版os.ReadFile()
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err = %v", err)
	}
	fmt.Printf("%v", string(bytes))
}
