package test

import "fmt"

/*var a = "runoob"

var b string = "run"

var c bool*/

var x, y int
var (
	a int
	b bool
)

var c, d = 1, 2

var  e, f = 123, "hello"
func init() {
	fmt.Print("Hello World!")
}

func main() {
	/*fmt.Print(a,b,c)*/
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
