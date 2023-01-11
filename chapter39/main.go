package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// 自定义比较规则
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

// 交换
func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println("----------")

	// 排序后
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}
