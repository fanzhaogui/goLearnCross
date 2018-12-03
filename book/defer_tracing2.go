package main

import "fmt"

func t(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}

func FFa()  {
	defer un(t("a"))
	fmt.Println("in a")
}

func FFb()  {
	defer un(t("b"))
	fmt.Println("in b")
	FFa()
}

func main()  {
	FFb()
	// entering b
	// in b
	// entering a
	// in a
	// leaving a
	// leaving b
}