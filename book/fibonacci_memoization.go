package main

import (
	"fmt"
	"time"
)

const LIM  = 41

var fibs [LIM]uint64

func main()  {
	var result uint64 = 0
	start := time.Now()

	for i := 0; i < LIM; i++ {
		result = Fi(i)
		fmt.Printf("Fi(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

func Fi(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		fibs[n] = 1
	} else {
		fibs[n] = Fi(n-1) + Fi(n-2)
	}
	res = fibs[n]
	return
}
