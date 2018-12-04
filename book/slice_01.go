// 切片 slice
//
// 给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，
// 切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：
// 切片是一个 长度可变的数组。
// 对于 切片 s 来说该不等式永远成立：0 <= len(s) <= cap(s)
// 优点 因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用。
// 声明切片的格式是： var identifier []type（不需要说明长度）。
// 一个切片在未初始化之前默认为 nil，长度为 0。
// 切片的初始化格式是：var slice1 []type = arr1[start:end]。

package main

import "fmt"

func main()  {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included
	// load the array with integers 0,1,2,3,4,5
	for i := 0; i < len(arr1); i ++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// 将切片传递给函数
	var arr = [5]int{0,1,2,3,4}
	sum(arr[:])

}
// 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i ++ {
		s += a[i]
	}
	return s
}

