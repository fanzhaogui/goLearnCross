package main

import (
	"fmt"
	"strconv"
)

// 字符串 类型转换
//
// 与字符串相关的类型转换都是通过 strconv 包实现的
// 任何类型 T 转换为字符串总是成功的
// strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数。
// strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。

// 将字符串转换为其它类型 tp 并不总是可能的，可能会在运行时抛出错误
// strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
// strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。


func main() {


	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)

	an = an + 5

	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
