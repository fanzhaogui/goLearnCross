package main

import "fmt"

func main() {

	goto One
	fmt.Printf("跳过中间的代码块")
	One:
		fmt.Println("这是代码块One")

}