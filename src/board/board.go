package board

import (
	"github.com/pedro-git-projects/go-chess/src/pieces"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

// MAYBE TODO ? Change coordiantes from strings to runes

// Board represents the a chessboard
type Board struct {
	coordinateMatrix [8][8]utils.Coordinate
	state            map[utils.Coordinate]pieces.Piece
}

// New returns a pointer to a board
func New() *Board {
	var matrix [8][8]utils.Coordinate
	clc, _ := utils.NewCircularCoordList('a')
	row := 8

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			matrix[i][j] = utils.NewCoordinate(clc.CurrentValue(), row)
			clc.MoveToNext()
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

// Find returns true if the coordinate is in the matrix
// false otherwise
func (b Board) Find(c utils.Coordinate) bool {
	for row := 0; row < len(b.coordinateMatrix); row++ {
		for col := 0; col < len(b.coordinateMatrix); col++ {
			if b.coordinateMatrix[row][col].First == c.First &&
				b.coordinateMatrix[row][col].Second == c.Second {
				return true
			}
		}
	}
	return false
}

// FilterByCol returns the column corresponding
// to the letter coordinate passed
func (b Board) FilterByCol(col rune) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)
	for i := 0; i < len(b.coordinateMatrix); i++ {
		for j := 0; j < len(b.coordinateMatrix); j++ {
			if b.coordinateMatrix[i][j].First == col {
				result = append(result, b.coordinateMatrix[i][j])
			}
		}
	}
	return result
}

// FilterByCol returns the column corresponding
// to the number coordinate passed
func (b Board) FilterByRow(row int) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)
	for i := 0; i < len(b.coordinateMatrix); i++ {
		for j := 0; j < len(b.coordinateMatrix); j++ {
			if b.coordinateMatrix[i][j].Second == row {
				result = append(result, b.coordinateMatrix[i][j])
			}
		}
	}
	return result
}

// FowardRightDiagonal returns the foward right diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) FowardRightDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second + 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x++
		y.MoveToNext()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// FowardLeftDiagonal returns the foward left diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) FowardLeftDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second + 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x++
		y.MoveToPrev()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// BackwardLeftDiagonal returns the backward left diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) BackwardLeftDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second - 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x--
		y.MoveToPrev()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// BackwardRightDiagonal  the backward right diagonal of a given
// board position as a possibly empty slice of coordinates
func (b Board) BackwardRightDiagonal(position utils.Coordinate) []utils.Coordinate {
	result := make([]utils.Coordinate, 0)

	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second - 1
	nextDiagonal := utils.NewCoordinate(y.CurrentValue(), x)

	for b.Find(nextDiagonal) {
		result = append(result, nextDiagonal)
		x--
		y.MoveToNext()
		nextDiagonal = utils.NewCoordinate(y.CurrentValue(), x)
	}

	return result
}

// FirstFoward returns the next square in front of the current one as well as a true value if it succeeds
// it returns an empty coordinate and not ok else
func (b Board) FirstFoward(position utils.Coordinate) (foward utils.Coordinate, ok bool) {
	next := position.Second + 1
	if next <= 8 {
		return utils.NewCoordinate(position.First, next), true
	}
	return utils.Coordinate{}, false
}

// NFoward returns the nth square in front of current as well as a true value if it succeeds
// it returns an empty coordinate and not ok else
func (b Board) NFoward(position utils.Coordinate, squares int) (foward utils.Coordinate, ok bool) {
	next := position.Second + squares
	if next <= 8 {
		return utils.NewCoordinate(position.First, next), true
	}
	return utils.Coordinate{}, false
}

// FowardRightL returns the position two squares foward one right from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) FowardRightL(position utils.Coordinate) (frl utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second + 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// FowardLeftL returns the position two squares foward one left from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) FowardLeftL(position utils.Coordinate) (fll utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second + 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// BackwardRightL returns the position two squares backward one right from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) BackwardRightL(position utils.Coordinate) (brl utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	x := position.Second - 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// BackwardLeftL returns the position two squares backward one left from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) BackwardLeftL(position utils.Coordinate) (brl utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	x := position.Second - 2
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// RightDownLateralL returns the position two squares right one down from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) RightDownLateralL(position utils.Coordinate) (rdll utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	y.MoveToNext()
	x := position.Second - 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// LeftDownLateralL returns the position two squares left one down from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) LeftDownLateralL(position utils.Coordinate) (ldll utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	y.MoveToPrev()
	x := position.Second - 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// LeftUpLateralL returns the position two squares left one up from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) LeftUpLateralL(position utils.Coordinate) (lull utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToPrev()
	y.MoveToPrev()
	x := position.Second + 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}

// RightUpLateralL returns the position two squares right one up from the current
// as well as an ok value if it succeeds
// else it returns a zeroed coordinate and false
func (b Board) RightUpLateralL(position utils.Coordinate) (rull utils.Coordinate, ok bool) {
	y, _ := utils.NewCoordList(position.First)
	y.MoveToNext()
	y.MoveToNext()
	x := position.Second + 1
	nextL := utils.NewCoordinate(y.CurrentValue(), x)

	if b.Find(nextL) {
		return nextL, true
	}
	return utils.Coordinate{}, false
}
