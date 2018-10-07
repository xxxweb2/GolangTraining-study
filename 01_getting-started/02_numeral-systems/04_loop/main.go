package main

import "fmt"

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Printf("%d \t %b \t %#x \n", i, i, i)
	}
}
