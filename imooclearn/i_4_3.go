package main

import (
	"fmt"
)

/* 控制语气

	条件语句 if , if...else ，嵌套if,else,if
	选择语句 switch, select
	循环语句for
	控制语句中使用到的关键字goto，break continue

*/

func main()  {
	a := 3
	if a > 0 {
		fmt.Println("a 大于 零")
		if a < 4 {
			fmt.Println("a 小于 4")
		}
	} else {
		fmt.Println("a 小于 零")
	}

	switch 3 {
	case 1:
		fmt.Println("值为1")
	case 2:
		fmt.Println("值为2")
	default:
		fmt.Println("以上都不是")

	}

	var b interface{}
	b = 32
	switch b.(type) {
	case int:
		fmt.Println("类型为int")
	case string:
		fmt.Println("类型为string")
	default:
		fmt.Println("以上类型都不是")



	}
}