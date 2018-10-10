package test

import "fmt"

/*
常量
const identifier [type] = value

iota 特殊常量， 可以认为是一个可以被编译器修改的常量
*/

const (
	d = "d"
	e = "e"
	f = "f"
)

func main()  {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println(LENGTH)
	fmt.Println(area)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
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
