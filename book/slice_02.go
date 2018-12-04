package main

import "fmt"

// 用make()创建一个切片
//
// 当相关数组还没有定义时，我们可以使用make()函数来创建一个切片，同时创建好相关数组
// var slice1 []type = make([]type, len)
// slice1 := make([]type, len)
// 这里 len 是数组的长度并且也是slice的初始长度
//
// make 接受2个参数：元素的类型以及切片的元素个数

func main()  {
	var slice1 []int = make([]int, 10)
	// load the array/slice
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\n The lenght of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}













