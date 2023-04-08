package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
	//skill string
}

func main() {
	// 将json字符串反序列化为struct
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Skill\":\"攻击\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败，err = %v\n", err)
	} else {
		fmt.Printf("反序列化成功，Name = %v,Age = %v,Skill = %v\n", monster.Name, monster.Age, monster.Skill)
	}

	// 反序列化map
	str02 := "{\"age\":90,\"name\":\"测试\"}"
	var a map[string]interface{}
	err02 := json.Unmarshal([]byte(str02), &a)
	if err02 != nil {
		fmt.Printf("反序列化失败，err = %v\n", err02)
	} else {
		fmt.Printf("反序列化成功，a = %v\n", a)
	}

	// 反序列化切片
	str03 := "[{\"age\":22,\"name\":\"测试2\"},{\"age\":29,\"name\":\"测试3\"}]"
	var slice []map[string]interface{}
	err03 := json.Unmarshal([]byte(str03), &slice)
	if err03 != nil {
		fmt.Printf("反序列化失败，err = %v\n", err03)
	} else {
		fmt.Printf("反序列化成功，slice = %v\n", slice)
	}
}
