package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://localhost:12345/people"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic d2FuZ3NodWJvOndhbmdzaHVibw==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "18774413-0c11-e312-7ed6-7bc4f8151f5a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
