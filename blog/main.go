package main

import (
	"net/http"
)

func main()  {
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}