package main

import (
	"fmt"
	"sort"
)

func main() {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔"
		monsters[1]["age"] = "200"
	}

	fmt.Println(monsters)

	newMonsters := map[string]string{
		"name": "new M",
		"age":  "100",
	}

	monsters = append(monsters, newMonsters)

	fmt.Println(monsters)

	// 排序
	map1 := make(map[int]int)
	map1[10] = 100
	map1[4] = 13
	map1[1] = 56
	map1[8] = 90

	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}

}
