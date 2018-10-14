package main

import (
	"github.com/gorilla/sessions"
	"net/http"
	"io"
	"fmt"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "session")
		if request.FormValue("email") != "" {
			// check password
			session.Values["email"] = request.FormValue("email")
		}
		session.Save(request, res)

		io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <form method="POST">
    `+ fmt.Sprint(session.Values["email"])+ `
      <input type="email" name="email">
      <input type="password" name="password">
      <input type="submit">
    </form>
  </body>
</html>`)
	})
	http.ListenAndServe(":9000", nil)
}
