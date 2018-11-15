package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame()  {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate()  {
	m.a = 5
	m.b = 7
}

/** 那么什么时候 reveiver 用指针呢
1. 改变receiver的值
2. struct本身非常的大，这样拷贝的代价是很昂贵的
3. 如果struct的一个method中receiver为指针，那么其他method的receiver最好也要用指针

method belongs to instance
function is a global function belongs to package
// 方法属于实例，函数是属于包的全局函数。

 */

// method 继承 如下
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human // 匿名字段
	school string
}

type Employee struct {
	Human // 匿名字段
	company string
}

// 在human上定义了一个method
func (h *Human) SayHi () {
	fmt.Printf("my name's %s ,you can call %s \n", h.name, h.phone)
}

// method 重写 -- method_function2
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s\n", e.name,
	e.company, e.phone)// split
}

func main()  {
	m := &Mutatable{0,0}
	fmt.Println(m)
	m.StayTheSame()
	fmt.Println(m)

	m.Mutate()
	fmt.Println(m) // value changed

	mark := Student{Human{"David", 23 ,"13655554444"}, "MIT"}
	sam := Employee{Human{"Messi", 33, "13722223333"},"Xueba"}

	mark.SayHi()
	sam.SayHi()
}