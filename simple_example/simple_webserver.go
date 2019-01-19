package main

import (
	"net/http"
	"io"
)

const form  = `

	<html><body>
	<form action="#" method="post" name="bar">
		<input name="in" type="test" />
		<input type="submit" value="submit">
	</form>
	</html></body>

`

func SimpleServer(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<h1>Hello, World</h1>")
}

func FormServer(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, r.FormValue("in"))
	}
}

func main()  {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}