package main

import (
	"fmt"
	"strings"
)

func main () {
	//var str string = "This is a example string"
	//fmt.Printf("T/F? Dose the string \"%s\" have prefix %s?", str, "Th")
	//fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	//strings.Contains(s, substr string ) bool  包含关系

	// 拼接 slice 到字符串

	str := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", s1)

	for _, val := range s1 {
		fmt.Printf("%s -", val)
	}

	fmt.Println()

	str2 := "GO1|The ABC of Go|25"
	s12 := strings.Split(str2, "|")

	for _, val := range s12 {
		fmt.Printf("%s -", val)
	}
	fmt.Printf("Splitted in slice: %v\n", s12)
	str3 := strings.Join(s12, ";")
	fmt.Printf("s12 joined by ;: %s\n", str3)
}
