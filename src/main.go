package main

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/board"
)

func main() {
	b := board.New()
	fmt.Println(b.State())
	fmt.Printf("_____________________________________\n")
	fmt.Println(b.StateStr())

}
