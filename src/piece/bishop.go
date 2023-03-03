package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Bishop type represents the bishop piece in a game of chess
type Bishop struct {
	color      color
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewBishop returns a new king instance
func NewBishop(color color, position utils.Coordinate) *Bishop {
	return &Bishop{
		color:    color,
		position: position,
	}
}

// Color is an accessor for the color field
func (b Bishop) Color() color {
	return b.color
}

// Position returns the current piece position
func (b Bishop) Position() utils.Coordinate {
	return b.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (b Bishop) LegalMoves() []utils.Coordinate {
	return b.legalMoves
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (b *Bishop) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)

	if b.color == White {

	}

	if b.color == Black {
	}

	b.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (b *Bishop) Move(to utils.Coordinate, board board) bool {
	b.CalculateLegalMoves(board)
	if utils.Contains(b.legalMoves, to) {
		b.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (b Bishop) String() string {
	return fmt.Sprintf("%s bishop", b.color)
}
