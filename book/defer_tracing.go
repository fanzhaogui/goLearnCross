package main

import "fmt"

func trace(s string)  {
	fmt.Println("entering:", s)
}
func untrace(s string)  {
	fmt.Println("leaving:", s)
}

func Fa()  {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func Fb()  {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	Fa()
}

func main()  {
	Fb()
	// entering b
	// in b
	// entering a
	// in a
	// leaving a
	// leaving b
}