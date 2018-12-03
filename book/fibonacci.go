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
	result := 0
	for i:=0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}

