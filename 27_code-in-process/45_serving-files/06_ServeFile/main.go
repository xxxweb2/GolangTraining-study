package main

import "net/http"

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}

func main() {
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":9000", nil)
}
