package main

import "fmt"

func main()  {
	str := ""
	for i := 1; i <= 25; i ++ {
		str += "G"
		fmt.Println(str)
	}

	for i := 1; i <= 9; i ++ {
		for j := 1; j <= 9; j ++ {
			if j <= i {
				fmt.Printf("%d*%d\t",i,j)
			}
		}
		fmt.Println()
	}
}
