package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 缓冲读取文件
func main() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}

	fmt.Printf("file=%v \n", file)
	fmt.Printf("file=%v \n", file.Name())

	// 默认缓冲区4096
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')

		fmt.Println(readString)

		if err == io.EOF { // 文件末尾
			break
		} // 读到一个换行就结束
	}

	err = file.Close()
	if err != nil {
		fmt.Println("close file err =", err)
	}
}
