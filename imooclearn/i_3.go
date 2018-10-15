package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 数据类型
// 整形 浮点型 复数 其他类型 布尔  派生

// 类型别名 type sz int32

type sz int32

func main()  {
	var a sz
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(unsafe.Sizeof(a))
}