// 匿名字段和内嵌结构体
//
// 结构体可以包含一个或多个匿名或内嵌字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是
// 字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体

package main

import "fmt"

type innerS struct {
	 in1 int
	 in2 int
}

type outerS struct {
	b int
	c float64
	int // anonymous field
	innerS // anonymous field
}

func main()  {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Println(outer)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println(outer2)


}
