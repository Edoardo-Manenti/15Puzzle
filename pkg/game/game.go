package game

import (
	"15puzzle/pkg/board"
	"15puzzle/pkg/coord"
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
	for {
		view.PrintBoard(b, terminal)
		var kp = view.ListenForInput(terminal)
		if kp == 4 {
			break
		}
		view.ClearTerminal(terminal)
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

