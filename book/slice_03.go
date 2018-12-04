// new() 和 make() 的区别
//
// new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为 T 的内存地址
//   这种方法  返回一个指向类型为 T，值为0的地址的指针，它适用于值类型如数组和结构体

// make(T) 返回一个类型为T的初始值，它只使用于3中内建的引用类型： 切片，map 和 channel

// 换言之，new 函数分配内存，make函数初始化

package main


