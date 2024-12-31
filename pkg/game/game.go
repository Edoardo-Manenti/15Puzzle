package game

import (
	"fmt"
	"15puzzle/pkg/board"
	//"15puzzle/pkg/coord"
	"15puzzle/pkg/view"
	"os"
	"io"
	"golang.org/x/term"
)


func Play(width int) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	b := board.NewBoard(width)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, "")
	input := make([]byte, 1)
	for {
		view.PrintBoard(b, terminal)
		fmt.Print("Select a move: (0:UP 1:RIGHT 2:DOWN 3:LEFT)->")
		_, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(input)
		//dir, _ := coord.GetDirection(in)
		//fmt.Printf("You selected the move %v\n", dir)
		//var err = b.Move(dir)
		//if err != nil {
		//	fmt.Println("The move you selected can not be executed")
		//	break
		//}
	}
}

func getInput() {

}

