package main

import "fmt"

func main()  {
	var ar [3]int
	af(ar) // [0,0,0]
	afp(&ar)// &[0,0,0]

	for i:=0; i < 3; i++ {
		afp(&[3]int{i, i*i, i*i*i})
	}
}

func af(a [3]int)  {
	fmt.Println(a)
}

func afp(a *[3]int)  {
	fmt.Println(a)
}
