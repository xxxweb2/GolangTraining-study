package main

import (
	"net/http"
	"strings"
	"io"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.RequestURI(), "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}

	io.WriteString(res, `
	Dog Name: <strong>`+ dogName+ `</strong><br>
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func main() {
	http.HandleFunc("/dog/", upTown)
	http.ListenAndServe(":9000", nil)
}
