package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["key"] = "123"
	fmt.Println(a)

	var b = make(map[string]string)
	b["key"] = "123"
	fmt.Println(b)

	var c = map[string]string{"key": "123"}
	fmt.Println(c)

}
