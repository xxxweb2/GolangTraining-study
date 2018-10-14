package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		key := "q"
		val := request.FormValue(key)
		fmt.Println("value:", val)
		res.Header().Set("Content-Type", "text/html;charset=utf-8")
		io.WriteString(res, `<form method="POST">

		 <input type="text" name="q">
		 <input type="submit">

		</form>`)
	})
	http.ListenAndServe(":9000", nil)
}
