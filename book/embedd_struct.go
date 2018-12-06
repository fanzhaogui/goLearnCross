// 内嵌结构体
//
// 同样地结构体也是一种数据类型，所以它也可以作为一个匿名字段来使用
// 外层结构体通过 outer.in1 直接进入内层结构体的字段，内嵌结构体甚至来自其他包
// 内层结构体被简单的插入或者内嵌进外层结构体
// 这个简单的 继承 机制提供了一种方式，使得可以从另外一个或一些类型继承部分或全部实现。

package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float64
}

func main()  {
	b := B{A{1,2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
