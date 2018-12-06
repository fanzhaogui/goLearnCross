// map 字典
//
// PHP中的关联数组
// var map1 map[kettype]valuetype
// var map1 map[string]int
//
// map 是 引用类型 的：内存用make方法来分配
// var map1 = make(map[keytype]valuetype)
// map1 : make(map[keytype]valuetype)

// 不要使用new， 永远用make来构造map
//  注意 如果你错误的使用 new()分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址

package main

import "fmt"

func main()  {
	var mapList map[string]int

	var mapAssigned map[string]int

	mapList = map[string]int{"one":1, "two":2}
	mapCreated := make(map[string]float64)
	mapAssigned = mapList

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assiged at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])


	mf := map[int]func() int {
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		},
	}

	fmt.Println(mf)
}
