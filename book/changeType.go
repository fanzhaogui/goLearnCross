package main

import "fmt"

// valueOfTypeB = typeB(valueOfTypeA)

type IZ int

var a IZ = 5

func main() {
	// 具有相同底层类型的变量之间可以相互转换
	c := int(a)
	d := IZ(c)
	fmt.Println(c)
	fmt.Println(d)
}

// 精度丢失： 当从一个取值范围较大的转换到取值范围较小的类型时（例如将int32转换为int16或将float32转换为int）

// 非法的类型转换，引发编译时错误