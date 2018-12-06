package main

import "fmt"

type File struct {
	fd int // 文件描述符
	name string
}

// 这个结构体类型对应的工厂方法
func NewFille(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main()  {
	file := NewFille(10, "./for_01.go")
	fmt.Print(file)
}