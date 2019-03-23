package main

import (
	"book/controllers"
	"fmt"
)

func main()  {
	fmt.Println("go的学习、熟练代码")

	i := controllers.NewIndex()
	i.Index()
}
