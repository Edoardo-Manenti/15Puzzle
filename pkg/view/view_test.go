package view


import (
	"testing"
	"15puzzle/pkg/board"
	"os"
	"io"
	"fmt"
	"golang.org/x/term"
)

func TestPrintBoard(t *testing.T) {
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, "")
	var b = board.NewBoard(4)
	PrintBoard(b, terminal)
}

func TestListenForInput(t *testing.T) {
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, "")
	var kp = ListenForInput(terminal)
	fmt.Println(kp)
}

func TestClearTerminal(t *testing.T) {
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, "")
	ClearTerminal(terminal)
}
