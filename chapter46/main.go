package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
	//skill string
}

func main() {

	// 结构体序列化
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

	// map序列化
	var a map[string]interface{}
	a = make(map[string]interface{})

	a["name"] = "测试"
	a["age"] = 90
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化为=%v\n", string(data))
	//a map 序列化为={"age":90,"name":"测试"}

	// 切片序列化
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "测试2"
	m1["age"] = 22
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "测试3"
	m2["age"] = 29
	slice = append(slice, m1)
	slice = append(slice, m2)
	data02, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice 序列化为=%v\n", string(data02))
	//slice 序列化为=[{"age":22,"name":"测试2"},{"age":29,"name":"测试3"}]

}
