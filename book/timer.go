package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()


	result := 0
	for i:=0; i <= 30; i++ {
		result = fibonacci3(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}


	end := time.Now()
	delta := end.Sub(start)
	fmt.Print("function took this amount of time: %s\n", delta)
}


func fibonacci3(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci3(n - 1) + fibonacci3(n - 2)
	}
	return
}

