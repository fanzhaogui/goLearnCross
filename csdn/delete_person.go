package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://localhost:12345/people/2"

	payload := strings.NewReader("{\n  \"firstname\": \"wang\",\n  \"lastname\": \"shubo\",\n  \"address\": {\n    \"city\": \"Beijing\",\n    \"state\": \"Beijng\"\n  }\n}")

	req, _ := http.NewRequest("DELETE", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "4c8d290e-4c6c-53f7-64e9-1d1f6ed19b09")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}