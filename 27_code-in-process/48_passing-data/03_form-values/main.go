package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		val := req.FormValue(key)
		fmt.Println("value: ", val)
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `<form method="GET">

		 <input type="text" name="q">
		 <input type="submit">

		</form>`+val)
	})
	http.ListenAndServe(":9000", nil)
}