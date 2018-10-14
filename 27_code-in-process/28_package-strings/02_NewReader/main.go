package main

import (
	"os"
	"log"
	"strings"
	"io"
)

func main() {
	dst, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	rdr := strings.NewReader("hello world")
	io.Copy(dst, rdr)
}
