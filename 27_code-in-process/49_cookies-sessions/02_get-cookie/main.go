package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		cookie, _ := request.Cookie("my-cookie")
		fmt.Println(cookie)
		fmt.Println(cookie.Value)

		http.SetCookie(res, &http.Cookie{
			Name:  "my-cookie",
			Value: "some value",
		})
	})
	http.ListenAndServe(":9000", nil)
}
