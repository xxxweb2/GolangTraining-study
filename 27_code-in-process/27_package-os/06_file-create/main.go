package main

import (
	"os"
	"log"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer f.Close()
}
