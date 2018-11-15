package main

import "fmt"

type MySlice []int
type EmptyObject struct {

}

var num3 = [5]int{1,2,3,4,5}
func main()  {
	fmt.Println(MySlice{})
	fmt.Println(EmptyObject{})

	for i,v := range num3 {
		fmt.Println(i,v)
	}
	var slice1 = num3[1:3] // 1 表示下界索引  3表示上界索引， 但不包含上界索引的值
	fmt.Println(num3[1])
	fmt.Println(num3[3])
	fmt.Println(slice1) // [2, 3]
	for i,v := range slice1 {
		fmt.Println(i,v)
	}

	fmt.Println("---------------")
	fmt.Println(slice1[0])// 2
	fmt.Println(len(slice1)) // 2
	fmt.Println(cap(slice1)) // 4  ==> 2-index=1 len=5  5-1=4

	fmt.Println("---------------")
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	slice3 := numbers3[2 : len(numbers3)]
	length := (3)
	capacity := (3)
	fmt.Println(slice3)// [3,4,5]
	fmt.Println(length,len(slice3))
	fmt.Println(capacity, cap(slice3)) // 3-index=2 len=5  5-2=3
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))

}
