package main

import (
	"fmt"
	"os"
)

// DumpFile prints file content to standard output
func DumpFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	fmt.Printf("content of %q:\n", filename)
	fmt.Println(string(data))
}
