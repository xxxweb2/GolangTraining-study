package main

import (
	"os"
	"fmt"
	"io"
)

func cp(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}

	io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v ", err)
	}
	return nil
}
