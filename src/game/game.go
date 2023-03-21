package game

import (
	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/piece"
)

var currentTurn turn

type Game struct {
	pieces []*piece.Piece
	board  *board.Board
}

func New() *Game {
	return &Game{
		board: board.New(),
	}
}
