package piece

import (
	"fmt"

	"github.com/pedro-git-projects/go-chess/src/utils"
)

// the Pawn type represents pawn pieces in a game of chess
type Rook struct {
	color      color
	moved      bool
	position   utils.Coordinate
	legalMoves []utils.Coordinate
}

// NewPawn returns a new pawn instance
func NewRook(color color, position utils.Coordinate) *Rook {
	return &Rook{
		color:    color,
		moved:    false,
		position: position,
	}
}

// Color is an accessor for the color field
func (r Rook) Color() color {
	return r.color
}

// Position returns the current piece position
func (r Rook) Position() utils.Coordinate {
	return r.position
}

// LegalMoves return the possibly empty list of legal moves
// for the piece given its current position
func (r Rook) LegalMoves() []utils.Coordinate {
	return r.legalMoves
}

// Moved sets the pawn moved field to true
func (r *Rook) Moved() {
	r.moved = true
}

// CalculateLegalMoves calculates the possibly empty
// slice of position the piece can move to
// given its current position
// and mutates the legalMoves field
func (r *Rook) CalculateLegalMoves(board board) {
	l := make([]utils.Coordinate, 0)
	col := board.FilterByCol(r.position.First)
	for _, pos := range col {
		for !board.IsOccupied(pos) {
			l = append(l, pos)
		}
	}
	row := board.FilterByRow(r.position.Second)
	for i, pos := range row {
		for !board.IsOccupied(pos) {
			l = append(l, pos)
		}
		l = append(l, row[i+1])
		break
	}

	r.legalMoves = l
}

// Move moves the piece to a given square
// if it is in the legal movement slice
func (r *Rook) Move(to utils.Coordinate, board board) bool {
	r.CalculateLegalMoves(board)
	if utils.Contains(r.legalMoves, to) {
		r.position = to
		return true
	}
	return false
}

// String retuns the piece name plus its color
func (r Rook) String() string {
	return fmt.Sprintf("%s rook", r.color)
}
