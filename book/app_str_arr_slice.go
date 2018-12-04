package main

import "fmt"

func main()  {
	s := "helloworld"
	c := []byte(s)
	fmt.Println(c)
	fmt.Println(s[1:5])
	fmt.Println(s[4:])
	fmt.Println(s[:4])
	fmt.Println(s[:len(s)])

	// 将字符串 hello 转换为 cello
	s = "hello"
	c = []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

}

// Compare 比较俩个字节数组字典顺序的整数对比
//
// == 0
// < -1
// > 1
func Compare(a, b[]byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch  {
		case a[i] < b[i]:
			return 1
		case a[i] > b[i]:
			return -1
		}
	}

	switch  {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
