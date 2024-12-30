package main

import (
	"fmt"
	"15puzzle/pkg/game"
//	"os"
//	"golang.org/x/term"
)

func main() {
	fmt.Println("Hello, this is the game '15Puzzle'")
	game.Play(4)
}

