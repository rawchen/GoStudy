package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func Teststruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	// 如果传参不是结构体直接返回
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取结构体字段数量
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		// 获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
		numofMethod := val.NumMethod()
		fmt.Printf("struct has %d methods\n", numofMethod)
		//var params []reflect.Value
		val.Method(1).Call(nil)

		//调用结构体的第1个方法Method(0)
		var params []reflect.Value
		params = append(params, reflect.ValueOf(10))
		params = append(params, reflect.ValueOf(40))
		res := val.Method(0).Call(params) // 传入的参数是[]reflect.Value
		fmt.Println("res=", res[0].Int()) // 返回结果是[]reflect.Value

	}
}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   20,
		Score: 99.9,
		Sex:   "男",
	}
	Teststruct(a)
}

//	struct has 4 fields
//	Field 0: 值为=黄鼠狼
//	Field 0: tag为=name
//	struct has 3 methods
//	---start---
//	{黄鼠狼 20 99.9 男}
//	---end---
//	res= 50
//	Field 1: 值为=20
//	Field 1: tag为=monster_age
//	struct has 3 methods
//	---start---
//	{黄鼠狼 20 99.9 男}
//	---end---
//	res= 50
//	Field 2: 值为=99.9
//	struct has 3 methods
//	---start---
//	{黄鼠狼 20 99.9 男}
//	---end---
//	res= 50
//	Field 3: 值为=男
//	struct has 3 methods
//	---start---
//	{黄鼠狼 20 99.9 男}
//	---end---
//	res= 50
