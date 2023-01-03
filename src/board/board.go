package board

import "github.com/pedro-git-projects/go-chess/src/utils"

type Board struct {
	CoordinateMatrix [8][8]utils.Pair[string, int]
}

func NewBoard() *Board {
	var matrix [8][8]utils.Pair[string, int]
	clc, _ := utils.NewCircularCoordList("a")
	row := 8

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			matrix[i][j] = utils.NewPair(clc.GetCurrent(), row)
			clc.Next()
		}
		row--
	}

	b := Board{
		CoordinateMatrix: matrix,
	}
	return &b
}
