package main

import (
	"math"
	"errors"
)

func main()  {

}


func MySqrt(f float64) (float64, error) {
	// return an error oas second parameter if invalid input
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	// otherwise use default square root function
	return math.Sqrt(f), nil
}


// name the number variable - by default it will have `zero-ed` values i.e. numbuers are 0, string is empty, etc
func MySqart2(f float64) (ret float64, err error)  {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(f)
	}
	return
}
