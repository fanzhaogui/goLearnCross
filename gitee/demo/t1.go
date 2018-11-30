// 题目： 古典问题： 有一队兔子，从出生后第3个月起，每个月都生一对兔子，
// 小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少?
package demo

import "fmt"

func Tfunc01(month int)  {
	f1, f2 := 1, 1
	// month := 6 //
	for i := 0; i < month; i++ {
		fmt.Printf("%d\n", f1)
		f1, f2 = f2, f1+f2
	}
}