package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 判断文件或目录是否存在
func main() {

	exists, err := PathExists("D:/test.txt")
	if err != nil {
		return
	}
	fmt.Println(exists)

	_, err = CopyFile("E:/test.txt", "D:/test.txt")
	if err != nil {
		fmt.Println("拷贝完成！")
	}
}

func PathExists(path string) (bool, error) {
	//path := "D:/test.txt"
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
