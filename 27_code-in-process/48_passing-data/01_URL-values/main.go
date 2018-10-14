package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		key := "q"
		val := request.URL.Query().Get(key)
		io.WriteString(writer, "Do my searchL:"+val)
	})
	http.ListenAndServe(":9000", nil)
}
