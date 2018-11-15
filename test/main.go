package main

import (
	"fmt"
	"log"
	"time"
)

const DATE_FORMAT  = "2018-11-11"
func main() {

	//var a int = 21
	//var b int = 10
	//var c int
	//
	//c = a + b
	//fmt.Print(c)
	//
	//c = a - b
	//fmt.Print(c)
	//
	//c = a * b
	//fmt.Print(c)
	//
	//c = a / b
	//fmt.Print(c)
	//
	//c = a % b
	//fmt.Print(c)
	//
	//a++
	//fmt.Print(a)
	//
	//a = 21
	//a--
	//fmt.Print(c)

	st := time.Now().Unix()
	et := time.Now().AddDate(0,0, 1).Unix()
	getDayStartUnix(st)
	getDayStartUnix(et)
}

func getDayStartUnix(sec int64) int64{
	t := time.Unix(sec, 0).Format("2006-01-02")
	sts, _ := time.Parse("2006-01-02", t)
	return sts.Unix()
}