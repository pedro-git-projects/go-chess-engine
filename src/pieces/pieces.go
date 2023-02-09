// pieces is a package for storing chess game pieces definitions
package pieces

import "github.com/pedro-git-projects/go-chess/src/utils"

// Piece is an interface with the restrictions a chess piece
// must satisfy
type Piece interface {
	Move(to utils.Coordinate)
	Color() Color
	Position() utils.Coordinate
	LegalMoves() []utils.Coordinate
	CalculateLegalMoves()
}
