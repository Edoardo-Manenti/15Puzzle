package view 

import (
	"fmt"
	"golang.org/x/term"
	"15puzzle/pkg/board"
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


func breakLine(sb *strings.Builder, w int, grid func(a ...interface{}) string) {
	(*sb).WriteString(grid("+---"))
	(*sb).WriteString(strings.Repeat(grid("+---"), w-2))
	(*sb).WriteString(grid("+---+"))
	(*sb).WriteString("\n")
}
