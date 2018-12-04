package main

import "fmt"

const(
	FIZZ  = 3
	BUZZ  = 5
	FIZZBUZZ = 15
)

func main() {

	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j ++ {
			print("G")
		}
		println()
	}

	// 按位补码从0到10
	for i := 1; i <= 10; i++ {

		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}

	for i := 1; i<= 100; i ++ {
		switch  {
		case i%FIZZBUZZ == 0:
			fmt.Println("FIZZBUZZ")
		case i%FIZZ == 0:
			fmt.Println("FIZZ")
		case i%BUZZ == 0:
			fmt.Println("BUZZ")
		}
	}
}
