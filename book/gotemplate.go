package main

// Go程序的执行（程序启动）顺序如下：
// 1 按照顺序导入所有被main包引用的其他包，然后在每个包中执行如下流程：
// 2 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次
// 3 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有init函数的话，则调用该函数。
// 4 在完成这一切之后，main也执行同样的过程，最后调用main函数开始执行程序。


import "fmt"

const c = "C"

var v int = 5

type T struct {
}

func init()  { // initialization of package

}

func main()  {
	var a int
	Func1()
	// ....
	fmt.Println(a)
}

func (t T) Method1() {
	// ...
}

func Func1() { // exported function Func1
	// ...
}
