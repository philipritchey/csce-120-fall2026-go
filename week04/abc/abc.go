package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	Main(os.Stdin, os.Stdout)
}

func Main(reader io.Reader, writer io.Writer) {
	n, slice, err := getInput(reader)
	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	i, j, k := abc(n, slice)

	printOutput(writer, slice, i, j, k)
}

func getInput(reader io.Reader) (int, []int, error) {
	scanner := bufio.NewScanner(reader)
	n, err := readInt(scanner)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid input, expected n to be an integer")
	}
	if n < 0 {
		return 0, nil, fmt.Errorf("invalid input, expected n >= 0")
	}

	slice, err := readInts(scanner, n)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid input, expected %d integers", n)
	}

	return n, slice, nil
}

func printOutput(writer io.Writer, slice []int, i, j, k int) {
	if i < 0 {
		fmt.Fprintln(writer, "impossible")
	} else {
		fmt.Fprintf(writer, "%d + %d = %d\n", slice[i], slice[j], slice[k])
		fmt.Fprintf(writer, "%d, %d, %d\n", i, j, k)
	}
}

// readInt reads a line of text and attempts to convert it to an int
func readInt(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("input is empty")
	}
	return strconv.Atoi(scanner.Text())
}

// readInts reads lines of text until it has read n ints or reaches EOF or an error
func readInts(scanner *bufio.Scanner, n int) ([]int, error) {
	slice := make([]int, n)
	i := 0
	for i < n && scanner.Scan() {
		s := scanner.Text()
		line := strings.NewReader(s)
		for i < n {
			_, err := fmt.Fscan(line, &slice[i])
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, fmt.Errorf("failed to read an int")
			}
			i++
		}
	}
	return slice, scanner.Err()
}

func abc(n int, slice []int) (int, int, int) {
	// handle "easy" cases
	if n < 3 {
		return -1, -1, -1
	}

	m := make(map[int]int)
	for i, v := range slice {
		m[v] = i
	}
	for i, a := range slice {
		for j, b := range slice {
			if j == i {
				continue
			}
			if k, ok := m[a+b]; ok && k != i && k != j {
				return i, j, k
			}

		}
	}
	return -1, -1, -1
}
