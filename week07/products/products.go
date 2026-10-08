package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int64
}

func main() {
	file, err := os.Open("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	_, err = csvReader.Read() // skip the header
	if err != nil {
		log.Fatal(err)
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var products []Product
	for _, record := range records {
		name := record[0]
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		quantity, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		product := Product{
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}
		products = append(products, product)
	}
	totalValue := 0.0
	for _, p := range products {
		value := p.Price * float64(p.Quantity)
		totalValue += value
		fmt.Printf("%s: $%0.2f, quantity: %d, value: $%0.2f\n", p.Name, p.Price, p.Quantity, value)
	}
	fmt.Printf("\nTotal inventory value: $%0.2f\n", totalValue)
}
