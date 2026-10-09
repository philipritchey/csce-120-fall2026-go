package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Song struct {
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	Seconds int    `json:"seconds"`
}

func main() {
	data, err := os.ReadFile("songs.json")
	if err != nil {
		log.Fatal(err)
	}
	var songs []Song
	err = json.Unmarshal(data, &songs)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	fmt.Println("All songs:")
	for _, song := range songs {
		fmt.Printf("%s by %s, %d seconds.\n", song.Title, song.Artist, song.Seconds)
		total += song.Seconds
	}

	fmt.Println()

	fmt.Println("Songs by Stack Trace:")
	for _, song := range songs {
		if song.Artist == "Stack Trace" {
			fmt.Println(song.Title)
		}
	}

	fmt.Println()

	fmt.Printf("Total playlist time: %d seconds\n", total)
}
