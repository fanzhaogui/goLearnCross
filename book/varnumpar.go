// 多个参数
//
// 如果变长参数的类型并不是都相同的呢？使用5个采参数来进行传递不是很明智的选择
package main

import "fmt"

func main()  {
	x := min(1, 3, 2, 0)
	fmt.Println(x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...) // 传递slice的方式
	fmt.Println(x)
}

func min(s ...int) int { // 多参数
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
