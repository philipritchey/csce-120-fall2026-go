package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("scores.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	_, err = r.Read() // skip the first (header) row
	s := 0.0
	cnt := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record[0], "scored", record[1])
		f, err := strconv.ParseFloat(record[1], 64)
		if err == nil {
			s += f
			cnt++
		}
	}
	fmt.Printf("Average score: %0.2f\n", s/float64(cnt))
}
