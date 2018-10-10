package test

import "fmt"

/*
运算符

	算术运算符
	关系运算符
	逻辑运算符
	位运算符
	复制运算符
	其他运算符
*/
/*
A = 0011 1100

B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001
*/

func main() {

	var a uint = 60    /* 60 = 0011 1100 */
	var b uint = 13    /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b       /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c )

	c = a | b       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c )

	c = a ^ b       /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c )

	c = a << 2     /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c )

	c = a >> 2     /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c )
}


/*
Go 程序可以由多个标记组成

关键字

	下面25个会使用到的关键字或保留字
 	break default func interface select
	case defer go map struct
	chan else goto pakage switch
	const fallthrough if range type
	continue for import return var

    36个预定义标识符
 	append bool byte cap close complex
	complex64 complex128 unit16 copy false float32
	float64 imag int int8 int16 unit32
	int32 int64 iota len make new
	nil panic uint64 print println real
	recover string true uint uint8 uintptr

	[] {} () .  ,  ;  ...

表示符

常量

字符串

符号

数据类型
1. 布尔型
2. 数字类型
3. 字符串类型
4. 派生类型

变量 var identifier type
 var 关键字
 变量名
 值
	1. 指定变量类型，声明后若不赋值，使用默认值
	var v_name v_type
	v_name = value

	2. 根据值来自行判定变量类型
	var v_name = value

	3. 省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误

*/
/*
多变量声明

	var vname1, vname2, vname3 type
	vname1, vname2, vname3 = v1, v2, v3

	var vname, vname2, vname3 = v1, v2, v3// 和python很像，不需要显示声明类型，自动推断
	vn1, vn2, vn3 = v1, v2, v3

	var {
		vname1 vtype1
		vname2 vtype2
	}
*/
/*
常量
const identifier [type] = value

iota 特殊常量， 可以认为是一个可以被编译器修改的常量
*/