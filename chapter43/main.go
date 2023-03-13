package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件，写入内容 5句“hello world”
	filePath := "d:/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 及时关闭file句柄
	defer file.Close()

	// 写入时使用带缓存的*Writer
	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 内容先写到缓存的，需要调用Flush真正写到文件中，否则文件中会没有数据
	writer.Flush()
}
