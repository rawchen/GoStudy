package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	//Num float64
	Num int
	//Name string
}

type Student struct {
	Name string
	Age  int
}

type Stu Student

type integer int

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
	//skill string
}

func main() {
	// 两结构体强转，要有相同的字段：字段个数、名称、类型要完全一致
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	var stu1 Student
	var stu2 Stu
	//stu2 = stu1  错误
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)

	var i integer = 10
	var j int = 20
	//j = i  错误
	j = int(i)
	fmt.Println(i, j)

	monster := Monster{"牛魔王", 500, ""}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("JSON序列化出错", jsonStr)
	}
	// [123 34 78 97 109 101 34 58 34 231 137 155 233 173 148 231 142 139 34 44 34 65 103 101 34 58 53 48 48 44 34 83 107 105 108 108 34 58 34 34 125]
	fmt.Println("jsonStr", jsonStr)
	//{"Name":"牛魔王","Age":500,"Skill":""}
	fmt.Println("jsonStr", string(jsonStr))

	// 如果结构体命名开头小写则为空串，因为别的包用到变量访问不到。
	// 给结构体添加TAG `json:"name"` 则可以使得大写开头的变量名序列化自定义。（json.Marshal()底层为反射）
	//{"name":"牛魔王","age":500,"skill":""}
}
