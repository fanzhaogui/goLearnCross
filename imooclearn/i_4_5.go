package main

import "fmt"

func main() {

	// 类似于for循环
	/*One:
		fmt.Println("这是代码块One")
		time.Sleep(1*time.Second)
	goto One*/

	/*for i:=1; i<4; i++ {
		if i >= 2 {
			break
		}
		fmt.Print("输出i值:")
		fmt.Println(i)
	}*/

	for i:=1; i<4; i++ {
		if i < 2 {
			continue
		}
		fmt.Print("输出i值:")
		fmt.Println(i)
	}

}