package main

import "fmt"

/*
函数
func function_name ([parameter list]) [return types] {
	todosomething
}

1. 函数的左花括号不能另起一行
2. 不允许函数内嵌定义
3. 支持多返回值，支持命名返回值
4. 函数只能判断是否为nil
5. 参数视为局部变量，因此不能声明同名变量
6. 不支持默认餐阿叔， _ 命名的参数也不能忽略
7. 支持可变参数
8. 可以在函数内定义匿名函数
9. 闭包
*/

// 1
func swap(x ,y string) (string, string) {
	return y, x
}

// 3
func split(sum int) (x, y int)  {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 6
func add(a, b int, _ bool) int {
	return a + b
}


// 7
func test(str string, a ...int)  {
	fmt.Println("%T, %v\n", str, a)
}

// 8


func main()  {
	fmt.Println(split(12)) // 5 7
	// fmt.Println(add(1,2)) // not enought arguments in call to add
	fmt.Println(add(1, 2, true))
	test("a", 1, 2, 3)

	func (s string) {
		fmt.Println(s)
	}("hello go!")
}
