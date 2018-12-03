// defer 和 追踪
//
// 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行  return 语句之后）一刻才执行某个语句或函数
// 为什么要在返回之后才执行这些语句？因为return语句同样可以包含一些操作，而不是单纯地返回某个值

// 使用defer的语句同样可以接收参数
// 当有多个defer行为被注册时，他们会以逆序执行（类似栈，即后进先出）

package main

import "fmt"

func main()  {
	function1()
	f() // 43210
}

func function1() {
	fmt.Printf("In function1 at the top \n")
	defer function2()
	//function2()
	fmt.Printf("In function2 at the bottom\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of calling function1\n")
}

func f() {
	for i := 0; i < 5; i++ { // 正常情况下打印 01234
		defer fmt.Printf("%d\n", i)
	}
}