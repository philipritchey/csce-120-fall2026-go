package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("message.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(data)
	fmt.Println(text)

	newText := "This file was created by my Go program.\n"
	err = os.WriteFile("output.txt", []byte(newText), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Wrote output.txt")
}
