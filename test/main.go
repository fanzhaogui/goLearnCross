package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Print(c)

	c = a - b
	fmt.Print(c)

	c = a * b
	fmt.Print(c)

	c = a / b
	fmt.Print(c)

	c = a % b
	fmt.Print(c)

	a++
	fmt.Print(a)

	a = 21
	a--
	fmt.Print(c)
}
