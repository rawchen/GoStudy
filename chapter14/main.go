package main

import "fmt"

var age int = 50
var Name string = "jack"

func test() {
	age := 10
	Name := "tom"
	fmt.Println("age =", age)
	fmt.Println("Name =", Name)
}

func main() {
	fmt.Println("age =", age)
	fmt.Println("Name =", Name)
	test()
}
