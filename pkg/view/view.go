package view 

import (
	"fmt"
	"golang.org/x/term"
	"15puzzle/pkg/board"
	"15puzzle/pkg/coord"
	"os"
	"strings"
	"github.com/fatih/color"
)

/*
** +---+---+
** | 3 |  4|
** +---+---+ 
** | 13| 14|
** +---+---+ 
** 
*/
func PrintBoard(b board.Board, terminal *term.Terminal) {
	var w = b.Width()
	var sb strings.Builder
	numbers := color.New(color.BgHiBlue, color.FgHiYellow, color.Bold).SprintFunc()
	grid := color.New(color.FgHiRed).SprintFunc()
	breakLine(&sb, w, grid)
	for _,r := range b.Data() {
		for _, v := range r {
			if v == 0 {
				sb.WriteString(fmt.Sprint(grid("|"), "   "))
			} else {
				sb.WriteString(grid("|"))
				sb.WriteString(numbers(fmt.Sprintf(" %2d", v)))
			}
		}
		sb.WriteString(grid("|"))
		sb.WriteString("\n")
		breakLine(&sb, w, grid)
	}
	fmt.Fprintln(terminal, sb.String())
}

func ListenForInput(terminal *term.Terminal) int, error {
	ansimap := map[[]byte]coord.Direction{
		[]byte{21,97,65}: Direction.UP,
		[]byte{21,97,66}: Direction.DOWN,
		[]byte{21,97,67}: Direction.RIGHT,
		[]byte{21,97,68}: Direction.LEFT,
	}
	color.Set(color.FgHiBlue, color.Underline, color.BlinkRapid)
	fmt.Print("To move press the arrow keys")
	color.Unset()
	input := make([]byte, 3)
	os.Stdin.Read(input)
	fmt.Println(input)
	return ansimap[input]
}

func ClearTerminal(terminal *term.Terminal) {
	fmt.Fprintln(terminal, "\033[H")
}

func breakLine(sb *strings.Builder, w int, grid func(a ...interface{}) string) {
	(*sb).WriteString(grid("+---"))
	(*sb).WriteString(strings.Repeat(grid("+---"), w-2))
	(*sb).WriteString(grid("+---+"))
	(*sb).WriteString("\n")
}
