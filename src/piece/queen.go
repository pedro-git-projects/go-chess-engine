package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Queen type represents the queen piece in a game of chess
type Queen struct {
	color      color
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewQueen returns a new queen instance
func NewQueen(color color, position utils.Coordinate) *Queen {
	return &Queen{
		color:    color,
		position: position,
	}
}

// Color is an accessor for the color field
func (q Queen) Color() color {
	return q.color
}

// Position returns the current piece position
func (q Queen) Position() utils.Coordinate {
	return q.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (q Queen) LegalMoves() []utils.Coordinate {
	return q.legalMoves
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (q *Queen) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)

	if q.color == White {

	}

	if q.color == Black {
	}

	q.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (q *Queen) Move(to utils.Coordinate, board board) bool {
	q.CalculateLegalMoves(board)
	if utils.Contains(q.legalMoves, to) {
		q.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (q Queen) String() string {
	return fmt.Sprintf("%s queen", q.color)
}
