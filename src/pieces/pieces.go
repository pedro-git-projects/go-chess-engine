// pieces is a package for storing chess game pieces definitions
package pieces

import "github.com/pedro-git-projects/go-chess/src/utils"

// Piece is an interface with the restrictions a chess piece
// must satisfy
type Piece interface {
	Move(to utils.Pair[string, int])
	Color() Color
	Position() utils.Pair[string, int]
	LegalMoves() []utils.Pair[string, int]
	CalculateLegalMoves()
}
