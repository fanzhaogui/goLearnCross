// 递归函数
//
// 经典案例： 斐波那契数列
// 即：前两个数为1，从第三个数开始每个数均为前两个数之和

package main

import (
	"fmt"
)

func main()  {
	//strings.Map("a", "asdfsadf")
	result := fbnc()
	for i:=0; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) is: %d\n", i, result())
	}
}

func fbnc() func() int {
	var a int = 0
	var b int = 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}

