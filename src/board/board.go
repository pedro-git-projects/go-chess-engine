package board

import (
	"github.com/pedro-git-projects/go-chess/src/pieces"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

// Board represents the a chessboard
type Board struct {
	coordinateMatrix [8][8]utils.Coordinate
	state            map[utils.Coordinate]pieces.Piece
}

// New returns a pointer to a board
func New() *Board {
	var matrix [8][8]utils.Coordinate
	clc, _ := utils.NewCircularCoordList("a")
	row := 8

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			matrix[i][j] = utils.NewCoordinate(clc.CurrentValue(), row)
			clc.Next()
		}
		row--
	}

	b := Board{
		coordinateMatrix: matrix,
	}
	return &b
}

// Matrix is an accessor for the board matrix
func (b Board) Matrix() [8][8]utils.Coordinate {
	return b.coordinateMatrix
}
