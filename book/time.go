package main

import (
	"time"
	"fmt"
)

// 时间操作文档： http://docs.studygolang.com/pkg/time/
//
// Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。Location 类型映射某个时区的时间，UTC 表示通用协调世界时间。

var week time.Duration

func main()  {
	t := time.Now()

	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
}

