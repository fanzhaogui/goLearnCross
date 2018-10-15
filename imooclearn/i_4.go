package main


/*
变量转换

	Go中不存在隐式转换，类型转换必须是显式的

    类型转换只能发生在两种兼容类型之间

    类型转换的格式

 		变量名称  [:] = 目标类型名称(需要转换的变量)

*/

/*

常量定义形式和常量类型范围

	显式： const identifier [type] = value
	隐式： const identifier = value (通常叫无类型常量)

特殊常量iota的使用

	iota在const关键字出现时将被充值为0 ；
	const(CONST1 CONST2 、、、) 中每新增一行常量声明将使得iota计算一次

    iota常见使用：
		1 跳值使用法
		2 插队使用法
		3 表达式隐式使用法
		4 单行使用法


常量可以使用内置表达式定义，例如：len() unsage.Sizeof()等

常量范围目前只支持布尔型、数值型（整数型、浮点型和复数）和字符串

*/


