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
	b.MovePiece(utils.Coordinate{'a', 2}, utils.Coordinate{'a', 4})
	ok := b.MovePiece(utils.Coordinate{'a', 2}, utils.Coordinate{'a', 4})
	fmt.Println(ok)
	b.MovePiece(utils.Coordinate{'b', 7}, utils.Coordinate{'b', 5})
	fmt.Println(b.StateStr())
	b.MovePiece(utils.Coordinate{'a', 4}, utils.Coordinate{'b', 5})
	fmt.Println(b.StateStr())
}
