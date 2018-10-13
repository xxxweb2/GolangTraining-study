package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
}
