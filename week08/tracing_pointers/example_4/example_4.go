package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

func main() {
	player := Player{Name: "Ava", Score: 10}

	p := &player
	p.Score = p.Score + 5

	fmt.Println(player.Score)
	fmt.Println(p.Score)
}
