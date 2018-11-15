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
}



