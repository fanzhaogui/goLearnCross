package main

import (
	"fmt"
	"imooclearn/show1"
	"imooclearn/show2"
)

func init()  {
	fmt.Println("main init")
}

func main() {
	fmt.Println("Hello world")
	show2.Show2()
	show1.Show1()

	var num4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := num4[4:6:8]
	fmt.Println(slice5, len(slice5))
}



