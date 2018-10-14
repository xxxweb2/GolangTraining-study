package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		key := "q"
		file, hdr, err := request.FormFile(key)
		fmt.Println(file, hdr, err)
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
      <input type="file" name="q">
      <input type="submit">
    </form>`)
	})

	http.ListenAndServe(":9000", nil)
}
