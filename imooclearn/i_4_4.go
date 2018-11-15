package main

import (
	"fmt"
	"time"
)

func main()  {
	/*for   { // 默认情况下为true
		fmt.Println("for times")
		time.Sleep(1*time.Second)
	}*/

	for i:=1; i<=10; i++ {
		fmt.Print("for times :")
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}

	a := []string{"apple", "banana", "egg"}
	for key,value:= range a { // 不限使用key的情况下可以使用 _ 替换
		fmt.Print("key的值为：")
		fmt.Println( key)
		fmt.Printf("value的值为：%s\n", value)
	}
}
