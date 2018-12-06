// map & struct vs new() & make()
//
package main

import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main()  {
	// ok
	y := new(Bar)
	(*y).thingOne = "Hello"
	(*y).thingTwo = 1
	fmt.Print(y)

	// NOT OK
	/*z := make(Bar) // 编译错误：cannot make type Bar
	(*z).thingOne = "hello"
	(*z).thingTwo = 1*/

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	//u := new(Foo)
	//(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	//(*u)["y"] = "world"
}
