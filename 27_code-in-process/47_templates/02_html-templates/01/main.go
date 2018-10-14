package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, Page{
		Title: "My title 2",
		Body:  `hello world <script>alert("Yow!");</script>`,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
