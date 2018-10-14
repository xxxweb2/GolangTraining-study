package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"io"
)

func main() {
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			src, _, err := request.FormFile("file")
			if err != nil {
				panic(err)
			}
			defer src.Close()

			dst, err := os.Create(filepath.Join("./", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()
			io.Copy(dst, src)
			err = tpl.Execute(res, nil)
			if err != nil {
				http.Error(res, err.Error(), 500)
				log.Fatalln(err)
			}
		}
	})

	http.ListenAndServe(":9000", nil)
}
