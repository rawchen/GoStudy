package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 创建一个可以存放任意类型的管道
	allChan := make(chan interface{}, 3)

	// 向管道写入数据
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{
		Name: "小花猫",
		Age:  20,
	}
	allChan <- cat

	// 推出2，读取第3个
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat = %T newCat = %v\n", newCat, newCat)

	// fmt.Printf("newCat.Name = %v\n", newCat.Name)
	// 不通过
	// 需要使用类型断言
	t := newCat.(Cat)
	fmt.Printf("t.Name = %v\n", t.Name)

}
