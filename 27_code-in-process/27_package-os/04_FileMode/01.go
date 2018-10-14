package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.ModeDir)
	fmt.Printf("%p\n", os.ModeDir)
	fmt.Printf("%d\n", os.ModeDir)
	fmt.Println(os.ModeAppend)
	fmt.Println(os.ModeExclusive)
	fmt.Println(os.ModePerm)
}
