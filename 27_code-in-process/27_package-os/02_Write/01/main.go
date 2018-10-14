package main

import (
	"os"
	"log"
)

func main() {

	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file %v", err)
	}
	defer dst.Close()
	dst.Write([]byte("Hello world"))
}
