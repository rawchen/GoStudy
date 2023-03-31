package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	// 遍历os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	// go main.go -u root
	var user string
	flag.StringVar(&user, "u", "", "用户名。默认为空")
	flag.Parse()

	fmt.Printf("user=%v", user)
}
