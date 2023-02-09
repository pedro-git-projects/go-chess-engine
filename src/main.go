package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
)

func main() {
	board := board.New()
	for i := 0; i < len(board.Matrix()); i++ {
		for j := 0; j < len(board.Matrix()); j++ {
			fmt.Printf("[%s%d]", board.Matrix()[i][j].First, board.Matrix()[i][j].Second)
		}
		println()
	}
}
