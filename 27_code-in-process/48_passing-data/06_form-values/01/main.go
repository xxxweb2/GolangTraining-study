package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		key := "q"
		file, _, err := request.FormFile(key)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		bs, _ := ioutil.ReadAll(file)
		fmt.Println(string(bs))
		res.Header().Set("Content-Type", "text/html")
		// you have to put this enctype for uploading files
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
      <input type="file" name="q">
      <input type="submit">
    </form>`)
	})
	http.ListenAndServe(":9000", nil)
}
