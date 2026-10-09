package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Programs often need to write data in memory out to a file
	// When they do this, the data has to get "marshalled" or "serialized"
	//   into some particular encoding that can be stored on disk
	// Text data is easy: string -> text/bytes
	// Numbers are easy: int, float -> text/bytes
	// But what about structs?
	// CSV is one common encoding: data as comma-separated text values
	//   field names -> header row (column names)
	//   field values -> row values
	// What if a field value is a struct?
	// CSV cannot easily handle this
	// But another common encoding can: JavaScript Object Notation (JSON)
	// JSON is very common in web programming
	// Like CSV, JSON data is text
	// Unlike CSV, JSON uses key-value pairs
	//   field names -> keys
	//   field values -> values
	// JSON object values can be nested quite naturally
	//   e.g. to handle structs that are composed of structs and arrays (and so on...)

	// This file contains an array/list of objects
	DumpFile("products.json")
	Pause()

	// As with CSV data, we want to associate a JSON object in a file with a Go struct in memory
	// Each key in a JSON object gets mapped to field name in a struct
	// Parsing JSON data is a non-trivial task
	// Luckily, there is the `encoding/json` package
	// (but you will one day soon be capable of writing a program that can parse JSON yourself)
	// In order to load JSON data into a struct using `encoding/json`,
	//   we tag the struct fields with their JSON key names

	type Product struct {
		Name     string  `json:"name"`
		Price    float32 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	// In order to be used by code outside our own package, the field names must be exported
	//   Exported names begin with a Capital letter
	// The names of the JSON keys in the tags should be camelCase (by convention)
	//   But MixedCase will also work (even with JSON files that use camelCase)

	// "Marshal" := convert data in memory into a format for storage or transmission
	// "Unmarshal" := the reverse process, convert stored or transmitted data back to in-memory structure

	// simulate the result of reading a file containing a single JSON object
	data := []byte(`{"name": "Notebook","price": 2.50,"quantity": 12}`)

	var product Product
	err := json.Unmarshal(data, &product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)

	// `json.Unmarshal(data, &product)`
	// means to load from `data` ([]bytes) into `product` (Product struct)

	Pause()

	// If the data contains a JSON array
	// Then unmarshal will expect to load into a slice of structs
	data, err = os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(products)

	Pause()

	// If the data and the struct do not align, an error will be returned
	// Try to unmarshal an array into a struct value
	err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println(err)
	}

	// An non-nil error will also happen in the case of invalid/malformed JSON syntax

	// Other cases will just fail silently:
	// * missing JSON key -> struct field gets "zero" value
	// * extra JSON key -> ignored

	Pause()

	// Last example: marshal data to JSON format
	product = Product{
		Name:     "Calculator",
		Price:    14.50,
		Quantity: 8,
	}
	data, err = json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	// data is []byte -> convert to string for human-readable
	fmt.Println(string(data))

	Pause()

	// data can be written to a file easily
	filename := "calculator.json"
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote", filename)

	Pause()
	DumpFile(filename)
}
