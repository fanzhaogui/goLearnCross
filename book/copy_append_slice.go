// 切片的复制和追加
//
// 拷贝切片的copy函数 和 向切片追加新元素的append函数

package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3, 4}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n)// n == len(sl_from) 复制了几个元素

	sl3 := []int{1,2,3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

