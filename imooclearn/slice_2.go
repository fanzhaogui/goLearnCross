package main

import "fmt"

func main()  {
	//var num4 = [...]int{1,2,3,4,5,6,7,8,9,10}
	//slice5 := num4[4:6:8] // 索引值   5-index=4   7-index=6   9-index=8
	//fmt.Println(num4[4], num4[6], num4[8]) // 5 7 9
	//fmt.Println(slice5) // [5, 6]
	//fmt.Println(len(slice5)) // 2
	//fmt.Println(cap(slice5)) // 10 - 6 = 4
	//slice5 = slice5[:cap(slice5)]
	//fmt.Println(slice5)// [5,6,7,8]
	//slice5 = append(slice5, 11, 12, 13)
	//fmt.Println(slice5)// [5,6,7,8,11,12,13]
	//slice6 := []int{0, 0, 0}
	//copy(slice5, slice6)// 把第二个参数值中的元素复制到第一个参数值中的相应位置（索引值相同）上
	//fmt.Println(slice5)// [0,0,0,8,11,12,13]
	//fmt.Println(slice5[2]) // 0
	//fmt.Println(slice5[3]) // 8
	//fmt.Println(slice5[4]) // 11

	//fmt.Println("----------------------")
	//var numbers3 = [5]int{1, 2, 3, 4, 5}
	//var slice1 = numbers3[1:4]
	//fmt.Println(slice1)
	//fmt.Println(cap(slice1)) // len 5 - indexfirst 1 = 4
	//slice1 = slice1[:cap(slice1)]
	//fmt.Println(slice1)
	//var slice2 = numbers3[1:4:4]
	//fmt.Println(slice2)
	//slice2 = slice2[:cap(slice2)]
	//fmt.Println(slice2)

	// append() copy()


	fmt.Println("------------------")

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8]
	length := (2)
	capacity := (4)
	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	slice5 = slice5[:cap(slice5)]
	fmt.Println(slice5)
	slice5 = append(slice5, 11, 12, 13)
	fmt.Println(slice5)
	length = len(slice5)// (7)
	fmt.Printf("%v\n", length)
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	fmt.Println(slice5)
	e2 := slice5[2]//(0)
	e3 := slice5[3]//(8)
	e4 := slice5[4]// (11)
	fmt.Printf("%v, %v, %v\n", e2, e3, e4)
}
