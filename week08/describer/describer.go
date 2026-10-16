package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

type Movie struct {
	Title    string
	Director string
}

type Song struct {
	Title  string
	Artist string
	Length int // seconds
}

// TODO: Define type Describer interface with one method: Describe() string
type Describer interface {
	Describe() string
}

// TODO: Implement Describe with a Book receiver
func (b Book) Describe() string {
	return b.Title + " by " + b.Author
}

// TODO: Implement Describe with a Movie receiver
func (m Movie) Describe() string {
	return m.Title + " directed by " + m.Director
}

func (s Song) Describe() string {
	return s.Title + " by " + s.Artist + " [" + strconv.Itoa(s.Length) + " seconds]"
}

func printDescription(item Describer) {
	fmt.Println(item.Describe())
}

// Expected output
// The Hobbit by J.R.R. Tolkien
// Spirited Away directed by Hayao Miyazaki

func main() {
	book := Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
	}

	movie := Movie{
		Title:    "Spirited Away",
		Director: "Hayao Miyazaki",
	}

	song := Song{
		Title:  "Ephemeral",
		Artist: "Intervals",
		Length: 225,
	}

	describables := []Describer{book, movie, song}
	for _, d := range describables {
		printDescription(d)
	}
}
