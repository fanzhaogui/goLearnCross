package main

import (
	"fmt"
)

type Act struct {
	Name string
	Age int
}

func main()  {
	var i1 = 5
	fmt.Println(i1, &i1)
	var intP *int
	intP = &i1
	fmt.Println(intP, *intP)

	var a1 Act
	a1.Name = "xiaoer"
	a1.Age = 20
	act := &a1
	fmt.Println(act, *act)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(s)
}
