// 方法 method
//
// 在Go中有这个概念，它和方法有着同样的名字，并且大体上意思相同：
// Go方法的作用在接收者receiver 上的一个函数，接收者是某种类型的变量。
// 因此方法是一种特殊类型的函数。

// 类型 T 或 *T 上的所有方法的集合叫做类型T(或*T)的方法集

// func (rec reveiver_type) methodName(parameter_list) (return_value_list) { ... }

// rec 就像面向对象语言中的 this 或 self，但是Go中并没有这两个关键字。随个人喜好，你可以使用
// this 或 self 作为 receiver 的名字

package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main()  {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is %d\n", two1.AddThem())
	fmt.Printf("The sum is %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is %d\n", two2.AddThem())

}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

