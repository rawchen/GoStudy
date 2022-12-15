package main

import "fmt"

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods  // 嵌套匿名结构体
	Writer string
}

func main() {
	var book = Book{
		Goods{
			Name:  "Tom",
			Price: 8,
		},
		"something",
	}

	fmt.Println(book)
}
