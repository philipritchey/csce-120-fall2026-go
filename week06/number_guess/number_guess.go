package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func main() {
	fmt.Println("You guess first.")
	HumanGuesser()

	fmt.Println("\nMy turn to guess.")
	ComputerGuesser()
}

func HumanGuesser() {
	fmt.Print("How big of a number can I pick? ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal("Not a number")
	}
	secret := rand.Intn(n) + 1
	guessesRemaining := int(math.Ceil(math.Log2(float64(n + 1))))
	fmt.Print(ColorText(Yellow, fmt.Sprintf("You have %d guesses.\n", guessesRemaining)))
	fmt.Println("Good luck!")
	guessNumber := 1
	for guessesRemaining > 0 {
		fmt.Printf("Guess #%d: ", guessNumber)
		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			log.Fatal("Not a number")
		}
		if guess < secret {
			fmt.Println(ColorText(Blue, fmt.Sprintf("%d is too low", guess)))
		} else if guess > secret {
			fmt.Println(ColorText(Red, fmt.Sprintf("%d is too high", guess)))
		} else {
			fmt.Println(ColorText(Green, "You got it!"))
			break
		}
		guessesRemaining--
		guessNumber++
	}
	if guessesRemaining == 0 {
		fmt.Println(ColorText(Yellow, "You are out of guesses. Better luck next time!"))
		fmt.Printf("My number was %d.\n", secret)
	}
}

func ComputerGuesser() {
	fmt.Println("Pick a secret number between 1 and 1000.")
	fmt.Println(ColorText(Yellow, "I will find it within 10 guesses."))
	low := 1
	high := 1000
	guessNumber := 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("Guess #%d is %s.\n", guessNumber, ColorText(Yellow, fmt.Sprint(mid)))
		var ans string
		fmt.Println("Tell me: my guess is...")
		fmt.Println("(l)ow")
		fmt.Println("(h)high")
		fmt.Println("(c)orrect")
		fmt.Print("? ")
		_, err := fmt.Scan(&ans)
		if err != nil {
			log.Fatal("failed to get input")
		}
		switch strings.ToLower(ans)[0] {
		case 'l':
			low = mid + 1
		case 'h':
			high = mid - 1
		case 'c':
			fmt.Println(ColorText(Green, "I got it!"))
			return
		}
		guessNumber++
	}
	fmt.Println(ColorText(Yellow, "I ran out of guesses... because you lied."))
}

func ColorText(color, s string) string {
	return color + s + Reset
}
