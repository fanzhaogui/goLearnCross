package main

import "fmt"

func main()  {
	typecheck(1, "12", []string{"a","b"})
}

func typecheck(values ... interface{})  {
	for _, value := range values {
		switch value.(type) {
		case int:
			//
		case float64:
			//
		default:
			//
		}
		fmt.Println(value)
	}

}