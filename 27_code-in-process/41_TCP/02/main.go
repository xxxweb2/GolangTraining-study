package main

import (
	"net"
	"io"
	"fmt"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, fmt.Sprint("hello world\n", time.Now()))
		conn.Close()
	}
}
