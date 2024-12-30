package view


import (
    "testing"
    "15puzzle/pkg/board"
)

func TestPrintBoard(t *testing.T) {
	var b = board.NewBoard(4)
	PrintBoard(b)
}
