package main

import (
	"net/http"
	"io"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}
func main() {
	var dog DogHandler
	var cat CatHandler
	mux := http.NewServeMux()
	mux.Handle("/dog", dog)
	mux.Handle("/cat", cat)
	http.ListenAndServe(":9000", mux)
}
