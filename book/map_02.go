// 测试键值对是否存在及删除元素
//
// 删除  delete(map1, key1) 如果key1不存在， 该操作不会产生错误

package main

import "fmt"

func main() {

	// 只是判断某个key是否存在而不关心它的值到底是多少
	// _, ok := map1[key1] // 如果key1存在则 ok == true,否则ok为false
	// if _, ok := map1[key1]; ok {
	//    // ....
	// }

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of Beijing in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is Paris in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is:%d\n", value)


	// delete an item
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of Washington in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Washington")
	}
}
