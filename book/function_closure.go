package main

import "fmt"

func main() {
	var f = Adder22()
	fmt.Print(f(1), "-") //  1
	fmt.Print(f(20), "-") // 21
	fmt.Print(f(300), "-") // 321
}

func Adder22() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
