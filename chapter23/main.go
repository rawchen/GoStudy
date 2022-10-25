package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["key"] = "123"
	fmt.Println(a)

}
