// Abstract Data Type || Record
//
// 结构体是复合类型 composite types ,当需要定义一个类型，它由一系列属性组成，每个属性都有自己的类型和值
// 的时候，就应该使用结构体，它把数据聚集在一起
package main

import (
	"book/libtest"
	"fmt"
)

type struct1 struct {
	i1 int
	f1 float64
	str string
}

type myStruct struct {i int}

func main()  {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	var v myStruct // v是结构体类型变量
	var p *myStruct // p是指向一个结构体类型变量的指针,初始值nil
	v.i = 2
	//p.i = 3 // 无法复制操作
	fmt.Print(v)
	fmt.Print(p)
	fmt.Println()
	// 初始化一个结构体实例（一个结构字面量）。更加简短和惯用的方式
	ms2 := &struct1{10, 15.5, "Chris"}
	// 或者
	var ms3 struct1
	ms3 = struct1{10, 15.5, "Chris"}
	fmt.Println(ms2,ms3)

	structExp := new(libtest.ExpStruct)

	structExp.Mi1 = 10
	structExp.Mf1 = 16.0


}

