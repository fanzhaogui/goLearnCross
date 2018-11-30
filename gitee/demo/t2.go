// 判断 101 - 200 之间有多少个素数，并输出所有素数
package demo

import (
	"math"
)

func Tfunc2() int {
	count := 1 // 合数的个数
	for i :=100; i <= 200; i++ {
		mid := int(math.Sqrt(float64(i)))
		for j := 2; j <= mid; j++ {
			if i % j == 0 {
				//fmt.Println(i) // 符合条件的即为合数
				count = count +  1
				break
			}
		}
	}
	return 100 - count
}
