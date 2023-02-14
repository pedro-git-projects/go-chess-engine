package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func main() {
	b := board.New()
	fmt.Println(b.State())
	pawn := b.PieceAt(utils.NewCoordinate('a', 2)) // refactor so as to not to have to retrieve a piece to move it
	pawn.Move(utils.NewCoordinate('a', 4), b)
	fmt.Println(b.State()) // all spaces need to be initialized in order for this func to work
}
