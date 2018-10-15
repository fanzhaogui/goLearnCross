package main

import (
	//imooc "fmt"
	. "fmt"
	_ "imooclearn/show2" // 只会执行 init
)

// import别名  . 和 _

// 别名操作： 将导入的包命名为另一个容易记忆的别名
// 点 操作： 点 . 标识的包导入后，调用该包中函数时可以忽略前缀包名
// 下划线 操作：导入该包，但不导入整个包，而是执行该包中的init函数，因此无法通过包名来调用包中的其他函数，
//		使用下划线操作往往是为了注册包里的引擎，让外部可以方便地使用

func main()  {
	//imooc.Println("alias pkg")
	Println("alias pkg")
	//Show2() // undefined: Show2

}