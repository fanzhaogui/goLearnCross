package main

import "fmt"

func main()  {
	//1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//2
	i := 2
START:
	fmt.Println(i)
	i++
	if i < 15 {
		goto START
	}

}
