package main

import (
	"fmt"
	"math"
)

// 区分 function 和 method 函数和方法

/*

1、通过func关键字声明方法

2、方法与函数的区别

方法和函数定义语法区别在于：
方法前置实例接受参数，这个receiver可以是基础类型也可以是指针。

3 receiver 传值 和 传指针的区别

*/

type Vertex struct {
	X, Y float64
}

// 1
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 2
type duck struct {

}

func (d *duck) quack()  { // receiver
	// do something
}
func quack(d *duck) { // function argument
    // do something
}

// 3 receiver 传值 和 传指针的区别

// 3.1 以下没有区别

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area_by_value() int {
	return r.length * r.width
}
func (r *Rectangle) Area_by_reference() int {
	return r.length * r.width
}

// 3.2 区别 ： 方法集
// 类型*T方法集合包含所有receiver T + *T 方法
type IFace interface {
	SetSomeField(newValue string)
	GetSomeField() string
}

type Implementation struct {
	someField string
}

// 获取一个值
func (i *Implementation) GetSomeField() string {
	return i.someField
}
// 设置一个值
func (i *Implementation) SetSomeField(newValue string) {
	i.someField = newValue
}

func Create() *Implementation {
	return &Implementation{someField:"Hello"}
}

func main()  {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	r1 := Rectangle{4, 3}
	fmt.Println(r1)
	fmt.Println("area is :", r1.Area_by_value())
	fmt.Println("area is :", r1.Area_by_reference())
	fmt.Println("area is :", (&r1).Area_by_value())
	fmt.Println("area is :", (&r1).Area_by_reference())

	var a IFace
	a = Create()
	fmt.Println(a.GetSomeField())
	a.SetSomeField("World")
	fmt.Println(a.GetSomeField())
}