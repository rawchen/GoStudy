package main

import "fmt"

func main() {
	var name string
	var n float32
	var isBool bool
	var text string
	var text2 string

	fmt.Scanln(&name)
	fmt.Println(name)

	fmt.Scanln(&n)
	fmt.Println(n)

	fmt.Scanln(&isBool)
	fmt.Println(isBool)

	fmt.Scanf("%v", &text)
	fmt.Println(text)

	fmt.Scanf("%v %v", &text2)
	fmt.Println(text2)

}
