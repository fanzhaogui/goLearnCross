package main

import "fmt"

func main()  {
	// 无限循环 递增
	//for i := 1 ;; i++ {
	//	fmt.Println(i)
	//}

	// 无限循环  000000000000000
	//for i := 0; i < 3 ; {
	//	fmt.Println(i)
	//}

	s := ""
	for ; s != "aaaaaa"; {
		fmt.Println("value of s:", s)
		s = s + "a"
	}

	for i,j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaaa"; i,j, s=i+1, j+1, s+"a" {
		fmt.Println("value of i, j, s:", i, j, s)
	}
}
