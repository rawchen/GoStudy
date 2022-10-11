package main

import (
	"fmt"
	"time"
)

func main() {
	// 日期时间相关函数
	now := time.Now()
	fmt.Println("now =", now)
	// 2022-10-11 22:45:12.5548283 +0800 CST m=+0.008197101
	fmt.Printf("type =%T \n", now)
	// time.Time
	fmt.Println(now.Year())   // 2022
	fmt.Println(now.Month())  // October
	fmt.Println(now.Day())    // 11
	fmt.Println(now.Hour())   // 22
	fmt.Println(now.Minute()) // 48
	fmt.Println(now.Second()) // 	1

	fmt.Println(now.UnixNano()) // 1665499769105985600	// 1970到现在为止的纳秒
	fmt.Println(now.Unix())     // 1665499769			// 1970到现在为止的秒数
	fmt.Println(now.Date())     // 2022 October 11
	fmt.Println(now.Clock())    // 22 49 29

	// 格式化
	fmt.Printf(now.Format("2006-01-02 15:04:05 \n")) // 2022-10-11 23:00:31
	fmt.Printf(now.Format("06 \n"))                  // 年 22
	fmt.Printf(now.Format("-07 \n"))                 // 时区 +08

	fmt.Println(time.Nanosecond)  // 1ns
	fmt.Println(time.Microsecond) // 1µs
	fmt.Println(time.Millisecond) // 1ms
	fmt.Println(time.Second)      // 1s
	fmt.Println(time.Minute)      // 1m0s
	fmt.Println(time.Hour)        // 1h0m0s

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond) // 休眠100毫秒
		fmt.Println("!!!:", i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(time.Now().Unix())
	}
}
