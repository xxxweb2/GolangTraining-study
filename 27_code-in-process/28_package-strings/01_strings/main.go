package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "tes"),
		strings.HasSuffix("test", "tes"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Replace("aaa", "a", "c", -1),
		strings.Split("a-c-v", "-"),
		strings.ToLower("TESt"),
		strings.ToUpper("TesT"),
		strings.TrimSpace("  Test  "),
	)

	arr := []byte("test")
	fmt.Println(arr)
}
