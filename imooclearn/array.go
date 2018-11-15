package main

import "fmt"

type MyNumbers [3]int

var numbers = [3]int{1, 2, 3}

func main()  {

	fmt.Println(numbers)
	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := (11)
	fmt.Println(numbers2)
	// “==”用于两个值的相等性判断
	fmt.Print(numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4])
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
}
