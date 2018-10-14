package main

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("session-id")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-id",
				Value: id.String() + " email=jon@email.com" + " JSON data" + " Whatever",
			}
			http.SetCookie(res, cookie)
		}
		fmt.Println(cookie)

	})
	http.ListenAndServe(":9000", nil)
}
