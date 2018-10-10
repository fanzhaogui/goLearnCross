// 程序所属包
package main

// 导入依赖包
import "fmt"

// 常量定义
const NAME string = "IMOOC"

// 全局变量的声明与赋值
var a string = "string 1"

// 一般类型赋值
var imoocInc int

// 结构的声明
type Learn struct {

}

//声明接口
type Ilearn interface {

}

// 函数定义
func learnImooc() {
	fmt.Println("learn imooc func")
}

// main()函数
func main () {
	learnImooc()
	fmt.Println("learn imooc")
}
