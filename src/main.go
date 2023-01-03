package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
)

func main() {
	board := board.NewBoard()
	for i := 0; i < len(board.CoordinateMatrix); i++ {
		for j := 0; j < len(board.CoordinateMatrix); j++ {
			fmt.Printf("[%s%d]", board.CoordinateMatrix[i][j].First, board.CoordinateMatrix[i][j].Second)
		}
		println()
	}
}
