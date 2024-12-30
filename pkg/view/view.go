package view 

import (
	"fmt"
	"15puzzle/pkg/board"
	"strings"
)

/*
** +---+---+
** | 3 |  4|
** +---+---+ 
** | 13| 14|
** +---+---+ 
** 
*/
func PrintBoard(b board.Board) {
	var w = b.Width()
	var sb strings.Builder
	breakLine(&sb, w)
	for _,r := range b.Data() {
		for _, v := range r {
			sb.WriteString(fmt.Sprintf("| %2d", v))
		}
		sb.WriteString("|\n")
		breakLine(&sb, w)
	}
	fmt.Println(sb.String())
}

func breakLine(sb *strings.Builder, w int) {
	(*sb).WriteString("+---")
	(*sb).WriteString(strings.Repeat("+---", w-2))
	(*sb).WriteString("+---+\n")
}
