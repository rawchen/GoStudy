package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["key"] = "123"
	fmt.Println(a)

	var b = make(map[string]string)
	b["key"] = "123"
	fmt.Println(b)

	var c = map[string]string{"key": "123"}
	fmt.Println(c)

	// crud
	c["key02"] = "456"
	c["key"] = "111"
	delete(c, "key")
	fmt.Println(c)

	val, ok := c["key02"]
	if ok {
		fmt.Println("c[\"no1\"]存在，值为=", val)
	}

	c = make(map[string]string)
	fmt.Println(c)

	// map遍历
	c["key01"] = "1"
	c["key02"] = "2"
	c["key03"] = "3"

	for k, v := range c {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
}
