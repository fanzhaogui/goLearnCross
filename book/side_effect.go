package main

import "fmt"

func main()  {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println(*reply)
}

// 改变外部变量
//
// 传递指针给函数，不但可以节省内存，因为没有复制变量的值，而且赋予了函数直接修改外部变量的能力
// 所以被修改的变量不再需要使用return返回，如下的例子，reply是指向 int 变量的指针，通过这个指针
// 我们在函数内，修改了这个int变量的数值
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
