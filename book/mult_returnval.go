package main

import "fmt"

func main()  {
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println(sum, prod, diff)
	sum, prod, diff = SumProductDiffn(3, 4)
	fmt.Println(sum, prod, diff)
}

func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func SumProductDiffn(i, j int) (s int, p int, d int) {
	s, p, d = i + j, i * j, i - j
	return
}