// 数组
//
// 具有相同 唯一类型 的一组已编号且长度固定的数据项序列
// 这种类型可以是任意的原始类型，例如：整形，字符串或者自定义类型
// 数组长度必须是一个常量表达式，并且必须是一个非负整数。
// 数组长度也是数组类型的一部分，所以[5]int 和 [10]int是属于不同类型的
// 数组编译时初始化是按照数组顺序完成的

package main

import "fmt"

func main()  {
	var arr1 [5]int
	for i:= 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i:= 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
}