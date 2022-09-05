package main

import (
	"GoStudy/chapter06"
	"fmt"
)

func main() {
	//26个字母 数组 下换线_
	//严格区分大小写
	//下划线为特殊的空标识符，对应的值会被忽略
	//变量名 函数名 常量名首字母大写则为公有public，其它包可访问，否则为小写本包访问private

	//key word

	//break		default		func	interface	select
	//case		defer		go		map			struct
	//chan		else		goto	package		switch
	//const 	fallthrouth if 		range 		type
	//continue 	for 		import	return		var

	var a int = 1
	fmt.Println(a)

	var _x int = 2
	fmt.Println(_x)

	var int int = 3
	fmt.Println(int)

	//引包测试
	chapter06.TestFunc()
	//chapter06.testFunc()
	fmt.Println(chapter06.A)
	//fmt.Println(chapter06.a)
}
