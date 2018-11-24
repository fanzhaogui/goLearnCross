package main

import (
	"fmt"
	"math"
)

// 如何定义
// 如何返回
// 方法和函数定义语法区别在于：

	// func (d *duck) quack() { // receiver
	//   do something
	// }


	// func quack(d *duck) { // function argument
	//  // do something
	// }

// 支持多返回值，支持命名返回值
func split2(sum int) (x, y int) {
	y = sum * 4 /9
	x = sum - x
	return
}

// 可变参数
func t2(str string, a ...int) {
	fmt.Println("%T, %v\n", str, a)
}

// 函数内定义匿名函数  可用余 异常或错误处理

// ** 区分函数和方法 **

type Vertest struct {
	a, b float64
}
// 声明方法
func (v Vertest) Abs() float64 {
	// 直接去使用定义类型中的变量
	return math.Sqrt(v.a * v.a + v.b * v.b)
}

type C struct { }
func (c *C)test()  {
	println("hi!", c)
}

type Cat struct {
	name string
	age int
}

func (c *Cat) growup()  { // 想想去掉 * 会怎么样
	c.age += 1
}

func (c Cat) getName() string {
	return c.name
}

func (c *Cat) setName(name string)  {
	c.name = name
}

// 加*  和  不加*的结果
/*

{wq 10} wq 10
加*
{wq 11} wq 11

{wq 10} wq 10
不加*
{wq 10} wq 10

*/




func main()  {
	fmt.Println(split2(17))
	t2("a", 1, 2, 3, 4)

	//err, code := func() (error, int) {
	//	b, err := strconv.Atoi("a")
	//	fmt.Println(b)
	//	return err, "there is no data"
	//}

	v := Vertest{3, 4}
	fmt.Println(v.Abs())

	var d C
	fmt.Println(d.test)


	cat := Cat{"wq", 10}
	fmt.Println(cat, cat.name, cat.age)
	cat.growup()
	fmt.Println(cat, cat.name, cat.age)

	cat.setName("gudqs")
	fmt.Println(cat, cat.name, cat.age)
	/*
		{wq 10} wq 10
		{wq 11} wq 11
		{gudqs 11} gudqs 11
	*/
}
