package main

import (
	"fmt"
	"os"
)

// 判断文件或目录是否存在
func main() {

	exists, err := PathExists("D:/test.txt")
	if err != nil {
		return
	}
	fmt.Println(exists)
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
