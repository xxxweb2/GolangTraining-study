package main

import (
	"net/http"
	"io"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world")
}

func main() {
	var h MyHandler
	http.ListenAndServe(":9000", h)
}
