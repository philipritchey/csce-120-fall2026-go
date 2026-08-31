package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// # get input
	// # convert to lowercase
	// # split on spaces
	// # join with underscores
	// # print
	// print('_'.join(input().lower().split()))
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(strings.Join(strings.Split(strings.ToLower(scanner.Text()), " "), "_"))
	}
}
