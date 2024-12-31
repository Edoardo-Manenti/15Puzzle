package view


import (
	"testing"
	"15puzzle/pkg/board"
	"os"
	"io"
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
