package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
)

func main() {
	// TODO: Encapsulate the piece movement methods inside the game pkg
	// so as not to need to pass pointers to board or create piece instances
	b := board.New()
	fmt.Println(b.StateStr())
}
