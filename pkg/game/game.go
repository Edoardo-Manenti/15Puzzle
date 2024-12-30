package game

import (
	"fmt"
	"15puzzle/pkg/board"
	"15puzzle/pkg/coord"
	"15puzzle/pkg/view"
)


func Play(width int) {
	b := board.NewBoard(width)
	var in int
	for true {
		view.PrintBoard(b)
		fmt.Print("Select a move: (0:UP 1:RIGHT 2:DOWN 3:LEFT)->")
		fmt.Scan(&in)
		dir, _ := coord.GetDirection(in)
		fmt.Printf("You selected the move %v\n", dir)
		var err = b.Move(dir)
		if err != nil {
			fmt.Println("The move you selected can not be executed")
			break
		}
	}
}
