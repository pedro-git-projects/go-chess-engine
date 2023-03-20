package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func main() {
	// TODO: Encapsulate the piece movement methods inside the game pkg
	// so as not to need to pass pointers to board or create piece instances
	b := board.New()
	fmt.Println(b.StateStr())
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	fmt.Println(b.StateStr())
	q := b.PieceAt(utils.NewCoordinate('d', 1))
	fmt.Println(q)

}
